package model

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"time"
	"net/http"
	"fmt"

	_ "github.com/lib/pq"
	"scooter/users/internal/processors"
)

var DB *gorm.DB

func InitDB(dataSourceName string) error {
    var err error

    DB, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
    if err != nil {
        return err
    }

	DB.AutoMigrate(&User{})

    return nil
}


type User struct {
	ID          uint `gorm:"unique;not null; uniqueIndex" json:"id"`
	Nome        *string `gorm:"not null" json:"nome"`
	CPF         *string `gorm:"unique;not null" json:"cpf"`
	Email       *string `gorm:"unique;not null" json:"email"`
	Celular 	*string `gorm:"not null" json:"celular"`
	CreatedAt   *time.Time
  	UpdatedAt   *time.Time
  }


func (u *User) ValidateUser() error {
    if u.Nome == nil {
        return fmt.Errorf("field 'nome' cannot be empty")
    }
    if u.CPF == nil {
        return fmt.Errorf("field 'cpf' cannot be empty")
    }
    if u.Email == nil {
        return fmt.Errorf("field 'email' cannot be empty")
    }
    if u.Celular == nil {
        return fmt.Errorf("field 'celular' cannot be empty")
    }
    return nil
}

func (u *User) UpdateUser(userUpdated *UserUpdate) {
    if userUpdated.Nome != nil {
        u.Nome = userUpdated.Nome
    }
    if userUpdated.Email != nil {
        u.Email = userUpdated.Email
    }
    if userUpdated.Celular != nil {
        u.Celular = userUpdated.Celular
    }
}

func (u *User) RaiseOnUserExists() error {
    userExists := User{}
	if err := DB.Where("cpf = ?", u.CPF).First(&userExists).Error; err == nil {
        // A user with the same CPF already exists, respond with an error
        return fmt.Errorf("A user with this CPF already exists")
    }
    return nil
}

func GetUserOr404(cpf string, w http.ResponseWriter, r *http.Request) *User {
	user := User{}
	if err := DB.First(&user, User{CPF: &cpf}).Error; err != nil {
		processors.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
