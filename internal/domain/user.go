package domain

import "time"

type User struct {
	ID        string
	FullName  string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
