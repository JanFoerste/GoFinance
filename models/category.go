package models

import (
	"gofinance/library"
	"time"
)

type Category struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCategoryInput struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"omitempty,hexcolor"`
}

type UpdateCategoryInput struct {
	Name  string             `json:"name"`
	Color library.JSONString `json:"color" binding:"omitempty,hexcolor"`
}
