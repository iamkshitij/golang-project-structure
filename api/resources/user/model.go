package user

import "time"

type User struct {
	ID        int       `db:"ID" json:"id"`
	Name      string    `db:"NAME" json:"name"`
	Email     string    `db:"EMAIL" json:"email"`
	Phone     string    `db:"PHONE" json:"phone"`
	Status    string    `db:"STATUS" json:"status"`
	CreatedAt time.Time `db:"CREATED_AT" json:"created_at"`
}
