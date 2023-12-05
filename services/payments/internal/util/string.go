package util

import (
    "errors"
    "github.com/google/uuid"
)



func StringToUUID(str string) (*uuid.UUID, error) {
    u, err := uuid.Parse(str)
    if err != nil {
        return nil, errors.New("Not valid UUID")
    }
    return &u, nil
}

func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}