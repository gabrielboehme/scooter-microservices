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