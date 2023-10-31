package config


import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var (
    EmployeeDB *gorm.DB
    DepartmentDB *gorm.DB
)

func ConnectEmployeeDB() {
    dsn := "root:system@tcp(localhost:3306)/empmanagement?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }
    EmployeeDB = db
}
func ConnectDepartmentDB() {
    dsn := "root:system@tcp(localhost:3306)/empmanagement?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the department database")
    }
    DepartmentDB = db
}
