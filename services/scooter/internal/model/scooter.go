package model

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"time"
	"net/http"
	"fmt"
    "errors"
    "github.com/google/uuid"
    
    "scooter/internal/processors"
    "scooter/internal/util"

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
    ID               *uuid.UUID `gorm:"unique;not null;default:gen_random_uuid()" json:"id"`
    SerialNumber     *uuid.UUID `gorm:"type:uuid;primaryKey; uniqueIndex; not null; default:gen_random_uuid()" json:"serial_number"`
    Latitude         *float64 `json:"latitude"`
    Longitude        *float64 `json:"longitude"`
    Status           *string  `json:"status"`
    State            *string `json:"state"`
	CreatedAt        *time.Time `json:"created_at"`
    UpdatedAt        *time.Time `json:"updated_at"`
}
var ScooterStatus = []string{"AVAILABLE", "IN_USE", "OUT_OF_ORDER"}
var ScooterStates = []string{"ON", "OFF"}

type Location struct {
    Latitude         *float64 `json:"latitude"`
    Longitude        *float64 `json:"longitude"`
}


func migrateDatabase(db *gorm.DB) {
    db.AutoMigrate(&Scooter{})
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

    if s.State == nil {
        return fmt.Errorf("field 'state' cannot be empty")
    }

    if s.ValidateScooterState() == false {
        return fmt.Errorf("Field 'state' must be one of %s", ScooterStates)
    }

    if s.ValidateScooterStatus() == false {
        return fmt.Errorf("Field 'status' must be one of %s", ScooterStatus)
    }

	scooterExists := Scooter{}
	if err := DB.Where("id = ?", s.ID).First(&scooterExists).Error; err == nil {
        // A scooter with the same id already exists, respond with an error
        return fmt.Errorf("A scooter with this ID already exists")
    }
    return nil
}

func (s *Scooter) ValidateScooterState() bool {
    return util.StringInSlice(
        *s.State,
        ScooterStates,
    )
}

func (s *Scooter) ValidateScooterStatus() bool {
    return util.StringInSlice(
        *s.Status,
        ScooterStatus,
    )
}

func (s *Scooter) UpdateScooter(scooterUpdated *ScooterUpdate) error {
    if scooterUpdated.State != nil {
        if util.StringInSlice(*scooterUpdated.Status, ScooterStatus) {
            s.Status = scooterUpdated.Status
        } else {
            return fmt.Errorf("Status must be one of %s", ScooterStatus)
        }
    }
    if scooterUpdated.State != nil {
        if util.StringInSlice(*scooterUpdated.State, ScooterStates) {
            s.State = scooterUpdated.State
        } else {
            return fmt.Errorf("State must be one of %s", ScooterStates)
        }
    }
    return nil
}

func (s *Scooter) ValidateLocationChange(newLocation *LocationUpdate) error {

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

    // Calculate the distance between the new location and the scooter's current location using the Haversine formula.
    distance := util.DistanceTwoLocations(*s.Latitude, *s.Longitude, *newLocation.Latitude, *newLocation.Longitude)

    // Compare the distance to 5km.
    if distance > 5 {
        return errors.New("The new location is more than 5km away from the scooter's current location")
    }

    s.Latitude = newLocation.Latitude
    s.Longitude = newLocation.Longitude

    return nil

}

func GetScooterOr404(serialNumber string, w http.ResponseWriter, r *http.Request) *Scooter {
	skooter := Scooter{}
    serialNumberUUID, err := util.StringToUUID(serialNumber)
    if err != nil {
        processors.RespondError(w, http.StatusNotFound, "serial_number is not a valid UUID.")
        return nil
    }
	if err := DB.First(&skooter, Scooter{SerialNumber: serialNumberUUID}).Error; err != nil {
		processors.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &skooter
}
