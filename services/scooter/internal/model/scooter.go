package model

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"time"
	"net/http"
	"fmt"
    "scooter/scooter/internal/processors"

    "github.com/google/uuid"

)

var DB *gorm.DB

func InitDB(dataSourceName string) error {
    var err error

    DB, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
    if err != nil {
        return err
    }

	DB.AutoMigrate(&Scooter{})

    return nil
}

type Scooter struct {
    ID               *uint `gorm:"unique;not null" json:"id"`
    SerialNumber     *uuid.UUID `gorm:"type:uuid;primaryKey; uniqueIndex; not null; default:gen_random_uuid()" json:"serial_number"`
    Latitude         *float64 `json:"latitude"`
    Longitude        *float64 `json:"longitude"`
    Status           *string  `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
  	UpdatedAt   *time.Time `json:"updated_at"`
}

type Location struct {
    Latitude         *float64 `json:"latitude"`
    Longitude        *float64 `json:"longitude"`
}

func migrateDatabase(db *gorm.DB) {
    db.AutoMigrate(&Scooter{})
    
    // Create constraint for status values -- update: doing in validateScooter
    // db.Exec("ALTER TABLE scooters ADD CONSTRAINT check_status CHECK (status IN ('Available', 'In Use', 'Under Repair'));")

}

func (s *Scooter) ValidateScooter() error {
    if s.ID != nil {
        return fmt.Errorf("cannot set 'id' field")
    }
    if s.Latitude == nil {
        return fmt.Errorf("field 'latitude' cannot be empty")
    }
    if s.Longitude == nil {
        return fmt.Errorf("field 'longitude' cannot be empty")
    }

    if s.Status == nil {
        return fmt.Errorf("field 'status' cannot be empty")
    }

	scooterExists := Scooter{}
	if err := DB.Where("id = ?", s.ID).First(&scooterExists).Error; err == nil {
        // A scooter with the same id already exists, respond with an error
        return fmt.Errorf("A user with this ID already exists")
    }
    return nil
}

func (s *Location) ValidateLocationChange(newLocation Location) error {

    if newLocation.Latitude == nil || newLocation.Longitude == nil {
        return errors.New("Both latitude and longitude must be provided")
    }

    // Add your custom validation logic here, such as bounds checking
    if *newLocation.Latitude < -90 || *newLocation.Latitude > 90 {
        return errors.New("Latitude is out of bounds")
    }

    if *newLocation.Longitude < -180 || *newLocation.Longitude > 180 {
        return errors.New("Longitude is out of bounds")
    }
    
    s.Location = newLocation

    return nil

}


func GetScooterOr404(serialNumber string, w http.ResponseWriter, r *http.Request) *Scooter {
	skooter := Scooter{}
	if err := DB.First(&skooter, Scooter{SerialNumber: &serialNumber}).Error; err != nil {
		processors.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &skooter
}
