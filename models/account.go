package models

import (
	"gofinance/library"
	"math/big"
	"time"
)

type Account struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Balance   big.Float `json:"balance"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateAccountInput struct {
	Name    string    `json:"name" binding:"required"`
	Balance big.Float `json:"balance"`
	Color   string    `json:"color" binding:"omitempty,hexcolor"`
}

type UpdateAccountInput struct {
	Name    string             `json:"name"`
	Balance big.Float          `json:"balance"`
	Color   library.JSONString `json:"color" binding:"omitempty,hexcolor"`
}
