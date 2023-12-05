package model

import (
    "encoding/json"
    "io"
)

type UserUpdateDecoder struct {
    Decoder *json.Decoder
}

type UserUpdate struct {
    Nome        *string `json:"nome"`
	Email       *string `json:"email"`
	Celular 	*string `json:"celular"`
}

func NewUserUpdateDecoder(r io.Reader) *UserUpdateDecoder {
    return &UserUpdateDecoder{Decoder: json.NewDecoder(r)}
}

func (d *UserUpdateDecoder) Decode(userUpdated *UserUpdate) error {
    if err := d.Decoder.Decode(userUpdated); err != nil {
        return err
    }
    return nil
}