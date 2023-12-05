package model

import (
    "net/http"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"time"
    "github.com/google/uuid"
	_ "github.com/lib/pq"
    "fmt"

    "rents/internal/processors"
    // "rents/internal/util"
)

var DB *gorm.DB

func InitDB(dataSourceName string) error {
    var err error

    DB, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
    if err != nil {
        return err
    }

	DB.AutoMigrate(&Rent{})

    return nil
}


type Rent struct {
	ID                      *uuid.UUID `gorm:"unique;not null;default:gen_random_uuid()" json:"id"`
    ScooterSerialNumber     *string
	RentStart               *time.Time `json:"rent_start"`
    RentFinish              *time.Time `json:"rent_finish"`
    CreatedAt               *time.Time `json:"created_at"`
  	UpdatedAt               *time.Time `json:"updated_at"`
  }

func GetInUseRentOr404(scooter string, w http.ResponseWriter, r *http.Request) *Rent {
    rent := Rent{}
	if err := DB.Where("scooter_serial_number = ?", scooter).First(&rent).Error; err != nil {
		processors.RespondError(w, http.StatusNotFound, fmt.Sprintf("Couldnt find rent for scooter %s", scooter))
		return nil
	}
	return &rent
  }