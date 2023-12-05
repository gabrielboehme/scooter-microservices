package model

import (
    "io"
    "encoding/json"
)

type ScooterUpdateDecoder struct {
    Decoder *json.Decoder
}

type ScooterUpdate struct {
    Status        *string `json:"status"`
	State       *string `json:"state"`
}

type LocationUpdateDecoder struct {
    Decoder *json.Decoder
}

type LocationUpdate struct {
    Latitude        *float64 `json:"latitude"`
	Longitude       *float64 `json:"longitude"`
}

func NewScooterUpdateDecoder(r io.Reader) *ScooterUpdateDecoder {
    return &ScooterUpdateDecoder{Decoder: json.NewDecoder(r)}
}

func NewLocationUpdateDecoder(r io.Reader) *LocationUpdateDecoder {
    return &LocationUpdateDecoder{Decoder: json.NewDecoder(r)}
}

func (d *ScooterUpdateDecoder) Decode(scooterUpdated *ScooterUpdate) error {
    if err := d.Decoder.Decode(scooterUpdated); err != nil {
        return err
    }
    return nil
}

func (d *LocationUpdateDecoder) Decode(locationUpdated *LocationUpdate) error {
    if err := d.Decoder.Decode(locationUpdated); err != nil {
        return err
    }
    return nil
}

func SerializeScooterToLocation(scooter *Scooter) Location {
    location := Location{
        Latitude:  scooter.Latitude,
        Longitude: scooter.Longitude,
    }
    return location
}