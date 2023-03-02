package models

import "time"

type Product struct {
	ProductID   string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Price       string    `gorm:"type:varchar(255)" json:"price"`
	Stock       string    `gorm:"type:varchar(255)" json:"stock"`
	Sold        string    `gorm:"type:varchar(255)" json:"sold"`
	Category    string    `gorm:"type:varchar(255)" json:"category"`
	Rating      float64   `gorm:"type:numeric" json:"rating"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp" json:"updated_at"`
}

type Products []Product
