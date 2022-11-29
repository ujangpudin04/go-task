package models

import "time"

type Ball struct {
	ID           int       `json:"id"`
	ClubHomeName string    `json:"clubhomename" gorm:"type: varchar(255)"`
	ClubWayName  string    `json:"clubwayname" gorm:"type: varchar(255)"`
	Score        string    `json:"score" gorm:"type: varchar(255)"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

type BallResponse struct {
	ID           int    `json:"id"`
	ClubHomeName string `json:"clubhomename"`
	ClubWayName  string `json:"clubwayname"`
	Score        string `json:"score"`
}

func (BallResponse) TableName() string {
	return "balls"
}
