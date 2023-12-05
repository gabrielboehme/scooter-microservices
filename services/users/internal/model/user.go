package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	db      *sql.DB
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	CPF     string `json:"cpf,omitempty"`
	Email   string `json:"email,omitempty"`
	Celular string `json:"celular,omitempty"`
}

func NewUserDB(db *sql.DB) *User {
	return &User{db: db}
}

func (c *User) Create(name string, cpf string, email string, celular string) (*User, error) {
	id := uuid.New().String()
	_, err := c.db.Exec(
		"INSERT INTO users (id, name, cpf, email, celular) VALUES ($1, $2, $3, $4, $5)",
		id, name, cpf, email, celular,
	)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:      id,
		Name:    name,
		CPF:     cpf,
		Email:   email,
		Celular: celular,
	}, nil
}

func (c *User) FindAll() ([]User, error) {
	rows, err := c.db.Query("SELECT id, name, cpf, email, celular FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []User{}
	for rows.Next() {
		var id, name, cpf, email, celular string
		if err := rows.Scan(&id, &name, &cpf, &email, &celular); err != nil {
			return nil, err
		}
		users = append(users, User{ID: id, Name: name, CPF: cpf, Email: email, Celular: celular})
	}
	return users, nil
}

func (c *User) Find(id string) (User, error) {
	var name, cpf, email, celular string
	err := c.db.QueryRow(
		"SELECT name, cpf, email, celular FROM users WHERE id = $1",
		id,
	).Scan(
		&name, &cpf, &email, &celular,
	)
	if err != nil {
		return User{}, err
	}
	return User{ID: id, Name: name, CPF: cpf, Email: email, Celular: celular}, nil
}

func (c *User) Update(name string, email string, celular string) (*User, error) {
	_, err := c.db.Exec(
		"UPDATE users SET name = $1, email = $2, celular $3",
		name, email, celular,
	)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:    name,
		Email:   email,
		Celular: celular,
	}, nil
}
