package controllers

import (
	"empManagement/config"
	"empManagement/models"
	"net/http"
	"github.com/gin-gonic/gin"
)


// Validate user
func ValidateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
 
    validatedUser, err := userService.ValidateUser(&user)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }
 
    c.JSON(http.StatusOK, validatedUser)
}