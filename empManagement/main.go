package main

import (
    "empManagement/config"
    "empManagement/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    config.ConnectEmployeeDB()
	config.ConnectDepartmentDB()
    routes.SetupRoutes(r)
    r.Run(":8081") // You can change the port as needed
}