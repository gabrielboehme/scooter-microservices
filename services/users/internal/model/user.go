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

type User struct {
	gorm.Model
	ID          uint `gorm:"unique;not null" json:"id"`
	Nome        *string `gorm:"not null" json:"nome"`
	CPF         *string `gorm:"unique;not null" json:"cpf"`
	Email       *string `gorm:"unique;not null" json:"email"`
	Celular 	*string `gorm:"not null" json:"celular"`
	CreatedAt   time.Time
  	UpdatedAt   time.Time
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

func GetUserOr404(cpf string, w http.ResponseWriter, r *http.Request) *User {
	user := User{}
	if err := DB.First(&user, User{CPF: &cpf}).Error; err != nil {
		processors.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}
