package models

type Department struct{
	ID     uint   `gorm:"primaryKey" json:"id"`
    Name   string `json:"name"`
}