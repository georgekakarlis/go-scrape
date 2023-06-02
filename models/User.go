package models

import "gorm.io/gorm"

// User struct
type User struct {
		gorm.Model
		FirstName string `gorm:"column:firstname;uniqueIndex;not null;size:50;" validate:"required,min=5,max=50" json:"firstname"`
		LastName  string `gorm:"column:lastname;uniqueIndex;not null;size:50;" validate:"required,min=5,max=50" json:"lastname"`
		Username  string `gorm:"uniqueIndex;not null;size:50;" validate:"required,min=3,max=50" json:"username"`
		Email     string `gorm:"uniqueIndex;not null;size:255;" validate:"required,email" json:"email"`
		Password  string `gorm:"not null;" validate:"required,min=6,max=50" json:"password"`
		Names     string `json:"names"`
	}
