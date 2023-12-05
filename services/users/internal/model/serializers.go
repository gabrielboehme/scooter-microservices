package model

import (
	// "fmt"
	"encoding/json"
)

type UserSerializer struct {
    User
}

type UserManySerializer struct {
	Users []SafeUser `json:"users"`
}

type SafeUser struct {
	Nome      string `json:"nome"`
	Email     string `json:"email"`
	Celular string `json:"celular"`
}


func NewUserSerializer(users []User) UserManySerializer {
    safeUsers := make([]SafeUser, len(users))
    for i, u := range users {
        safeUsers[i] = SafeUser{
            Nome:      *u.Nome,
            Email:     *u.Email,
            Celular: *u.Celular,
        }
    }
    return UserManySerializer{Users: safeUsers}
}

func (s UserSerializer) MarshalJSON() ([]byte, error) {
    safeUser := SafeUser{
        Nome:      *s.User.Nome,
        Email:     *s.User.Email,
        Celular:   *s.User.Celular,
    }

    // Marshal the SafeUser to JSON
    return json.Marshal(safeUser)
}