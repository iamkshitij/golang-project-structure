package user

import (
	"github.com/gofiber/fiber/v3/log"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

// GetAllUsers retrieves all users from the database ordered by creation date (newest first)
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllUsers() ([]User, error) {
	log.Debug("Executing GetAllUsers query")
	var users []User
	query := `SELECT ID, NAME, EMAIL, PHONE, STATUS, CREATED_AT FROM users ORDER BY CREATED_AT DESC`

	err := r.db.Select(&users, query)
	if err != nil {
		log.Errorf("Database error in GetAllUsers: %v", err)
		return nil, err
	}

	log.Debugf("Retrieved %d users from database", len(users))
	return users, nil
}

func (r *Repository) GetUserByID(id int) (*User, error) {
	log.Debugf("Executing GetUserByID query for ID: %d", id)
	var user User
	query := `SELECT ID, NAME, EMAIL, PHONE, STATUS, CREATED_AT FROM users WHERE ID = :1`

	err := r.db.Get(&user, query, id)
	if err != nil {
		log.Errorf("Database error in GetUserByID for ID %d: %v", id, err)
		return nil, err
	}

	log.Debugf("Retrieved user: %s (ID: %d)", user.Email, user.ID)
	return &user, nil
}

func (r *Repository) CreateUser(user *User) error {
	log.Debugf("Executing CreateUser query for: %s", user.Email)
	query := `INSERT INTO users (NAME, EMAIL, PHONE, STATUS) 
			  VALUES (:1, :2, :3, :4) 
			  RETURNING ID, CREATED_AT INTO :5, :6`

	_, err := r.db.Exec(query, user.Name, user.Email, user.Phone, "active", &user.ID, &user.CreatedAt)
	if err != nil {
		log.Errorf("Database error in CreateUser for %s: %v", user.Email, err)
		return err
	}

	log.Debugf("Created user with ID: %d", user.ID)
	return nil
}

func (r *Repository) UpdateUser(id int, user *User) error {
	log.Debugf("Executing UpdateUser query for ID: %d", id)
	query := `UPDATE users SET NAME = :1, EMAIL = :2, PHONE = :3, STATUS = :4 WHERE ID = :5`

	_, err := r.db.Exec(query, user.Name, user.Email, user.Phone, user.Status, id)
	if err != nil {
		log.Errorf("Database error in UpdateUser for ID %d: %v", id, err)
		return err
	}

	log.Debugf("Updated user with ID: %d", id)
	return nil
}

func (r *Repository) DeleteUser(id int) error {
	log.Debugf("Executing DeleteUser query for ID: %d", id)
	query := `DELETE FROM users WHERE ID = :1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Errorf("Database error in DeleteUser for ID %d: %v", id, err)
		return err
	}

	log.Debugf("Deleted user with ID: %d", id)
	return nil
}
