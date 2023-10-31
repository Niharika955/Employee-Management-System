package models


type Complaince struct{
	complianceId int `gorm:"primaryKey" json:"complianceid"`
	riType string `json: "ritype"`
	details string `json: "details"`
	createDate time.Time `json: "createdate"`
	department Department `json:"department"`
	empCount int `json: "empcount"`
	stsCount int `json: "stscount"`
	status string `json: "status"`
}