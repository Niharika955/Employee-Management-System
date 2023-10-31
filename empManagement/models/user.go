package models

type User struct{
	userId int `gorm:"primaryKey" json:"userid"`
	password string `json:"password"`
	role string `json:"role"`
	result int `json:"result"`
}

