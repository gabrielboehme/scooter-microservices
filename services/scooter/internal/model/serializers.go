package model

func SerializeScooterToLocation(scooter model.Scooter) Location {
    location := Location{
        Latitude:  scooter.Location.Latitude,
        Longitude: scooter.Location.Longitude,
    }
    return location
}