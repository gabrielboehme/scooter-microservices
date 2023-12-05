package model

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"time"
    "github.com/google/uuid"
	_ "github.com/lib/pq"
    "errors"
)

var DB *gorm.DB

func InitDB(dataSourceName string) error {
    var err error

    DB, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
    if err != nil {
        return err
    }

	DB.AutoMigrate(&Payment{})

    return nil
}


type Payment struct {
	ID           *uuid.UUID `gorm:"unique;not null;default:gen_random_uuid()" json:"id"`
	CardNumber   *string `json:"card_number"`
    RentId       *string `json:"rent_id"`
    CreatedAt    *time.Time `json:"created_at"`
  }

func (p *Payment) ValidatePayment() error {
    if p.ID != nil {
        return errors.New("Cant set payment ID")
    }
    if p.CardNumber == nil {
        return errors.New("Card number cant be null")
    }
    if p.RentId == nil {
        return errors.New("Card number cant be null")
    }
    return nil
}