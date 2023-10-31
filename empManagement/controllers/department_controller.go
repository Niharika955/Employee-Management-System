package controllers

import (
    "empManagement/config" // Replace with your actual project package name
    "empManagement/models"  // Replace with your actual project package name
    "github.com/gin-gonic/gin"
    "net/http"
)

// CreateDepartment creates a new department.
func CreateDepartment(c *gin.Context) {
    var department models.Department
    if err := c.ShouldBindJSON(&department); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DepartmentDB.Create(&department).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create department"})
        return
    }

    c.JSON(http.StatusCreated, department)
}

// GetDepartments retrieves a list of all departments.
func GetDepartments(c *gin.Context) {
    var departments []models.Department

    if err := config.DepartmentDB.Find(&departments).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch departments"})
        return
    }

    c.JSON(http.StatusOK, departments)
}

// GetDepartmentDetails retrieves details of a specific department by ID.
func GetDepartmentDetails(c *gin.Context) {
    id := c.Param("id")
    var department models.Department

    if err := config.DepartmentDB.Where("id = ?", id).First(&department).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
        return
    }

    c.JSON(http.StatusOK, department)
}

// DeleteDepartment deletes a department by ID.
func DeleteDepartment(c *gin.Context) {
    departmentID := c.Param("id")

    if err := config.DepartmentDB.Where("id = ?", departmentID).Delete(&models.Department{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete department"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Department deleted successfully"})
}
