package models

import (
	"time"

	"github.com/google/uuid"
)

// custom initialization of gorm.Model, we are calling it base
// can be used as standard for all structs that need the same default columns like ID as uuid
type Base struct {
	ID        uuid.UUID `gorm:"primary_key; unique; type:uuid; default:gen_random_uuid()"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// TODO: implement deposit, withdraws and total columns in database
// User object with some rules for SQL schema, generated by the "gorm:" tags
type User struct {
	Base
	Username     string        `validate:"required" gorm:"unique_index;not null" json:"username"`
	Email        string        `validate:"required,email" gorm:"unique_index;not null" json:"email"`
	Password     string        `validate:"required,min=8" gorm:"not null" json:"password"`
	Transactions []Transaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Transaction struct {
	Base
	Description string `json:"description" gorm:"not null" validate:"required"`
	Number      int    `json:"number" gorm:"not null" validate:"required,min=1"`
	Category    string `json:"category" gorm:"not null" validate:"required"`
	Type        string `json:"type" gorm:"not null" validate:"required"`
	UserID      uuid.UUID
}
