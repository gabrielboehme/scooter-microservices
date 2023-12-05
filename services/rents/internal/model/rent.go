package model

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"time"
    "github.com/google/uuid"

	_ "github.com/lib/pq"
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
