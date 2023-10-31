package models

type StatusReport struct{
	statusId int `gorm:"primaryKey" json:"statusid"`
	comment string `json:"comment"`
	details string `json:"details"`
	createDate time.Time `json:"createdate"`
	userId int `json:"userid"`
	complianceId int `json:"complianceid"`
	department Department `json:"department"`
}
