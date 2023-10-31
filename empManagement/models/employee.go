package models


type Employee struct {
    ID          uint   `gorm:"primaryKey" json:"id"`
    Name        string `json:"name"`
    Email       string `json:"email"`
    UserID      string `json:"user_id"`
    DepartmentID uint   `json:"department_id"`
    Department   Department `gorm:"foreignkey:DepartmentID"` // Define the relationship
}
