package user

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

type Resource struct {
	repo *Repository
}

func NewResource(db *sqlx.DB) *Resource {
	return &Resource{
		repo: NewRepository(db),
	}
}

func (r *Resource) GetAllUsers() ([]User, error) {
	return r.repo.GetAllUsers()
}

func (r *Resource) GetUserByID(id int) (*User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}
	
	return r.repo.GetUserByID(id)
}

func (r *Resource) CreateUser(user *User) error {
	// Basic validation
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	
	// Set default status if not provided
	if user.Status == "" {
		user.Status = "active"
	}
	
	return r.repo.CreateUser(user)
}

func (r *Resource) UpdateUser(id int, user *User) error {
	if id <= 0 {
		return errors.New("invalid user ID")
	}
	
	// Basic validation
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	
	return r.repo.UpdateUser(id, user)
}

func (r *Resource) DeleteUser(id int) error {
	if id <= 0 {
		return errors.New("invalid user ID")
	}
	
	return r.repo.DeleteUser(id)
}