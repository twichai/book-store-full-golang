package entities

import (
	"gorm.io/gorm"
)

type Publisher struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name,omitempty"`
	Address    string `json:"address,omitempty"`
	Phone      string `json:"phone,omitempty"`
}
