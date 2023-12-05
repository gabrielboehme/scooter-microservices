package model



func SerializeScooterToLocation(scooter *Scooter) Location {
    location := Location{
        Latitude:  scooter.Latitude,
        Longitude: scooter.Longitude,
    }
    return location
}