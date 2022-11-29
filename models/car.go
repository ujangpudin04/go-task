package models

import "time"

type Car struct {
	ID        int       `json:"id"`
	Price     string    `json:"price" gorm:"type: varchar(255)"`
	Brand     string    `json:"brand" gorm:"type: varchar(255)"`
	Type      string    `json:"type" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type CarResponse struct {
	ID    int    `json:"id"`
	Price string `json:"price"`
	Brand string `json:"brand"`
	Type  string `json:"type"`
}

func (CarResponse) TableName() string {
	return "cars"
}
