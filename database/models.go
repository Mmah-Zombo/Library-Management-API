package database

import "time"

type Book struct {
	ID          uint       `gorm:"primaryKey"`
	Title       string     `json:"title" form:"title"`
	Author      string     `json:"author" form:"author"`
	PublishDate *time.Time `json:"publish_date" form:"publish_date"`
	CreatedAt   time.Time  `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" form:"updated_at"`
}
