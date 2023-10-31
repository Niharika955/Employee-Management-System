package controllers

import (
	"empManagement/config"
	"empManagement/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create a new Compliance document (RL document)
func CreateCompliance(c *gin.Context) {
    var compliance models.Compliance
    if err := c.ShouldBindJSON(&compliance); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
 
    if err := complianceService.CreateRL(&compliance); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Compliance document"})
        return
    }
 
    c.JSON(http.StatusCreated, compliance)
}
 
// Get all Compliance documents
func GetComplianceDocuments(c *gin.Context) {
    complianceList := complianceService.GetAIRL()
    c.JSON(http.StatusOK, complianceList)
}
 
// Get Compliance documents by user ID
func GetComplianceByUser(c *gin.Context) {
    userld := c.Param("userld")
    complianceList := complianceService.GetAlIRL(userld)
    c.JSON(http.StatusOK, complianceList)
}
 
// Create a new Status Report
func CreateStatusReport(c *gin.Context) {
    var statusReport models.StatusReport
    if err := c.ShouldBindJSON(&statusReport); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
 
    if err := complianceService.CreateStatusReport(&statusReport); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Status Report"})
        return
    }
 
    c.JSON(http.StatusCreated, statusReport)
}
 
// Get all Status Reports by user ID and Compliance ID
func GetStatusReports(c *gin.Context) {
    userld := c.Param("userld")
    compid := c.Param("compid")
    statusReports := complianceService.GetAllStatusReport(userld, compid)
    c.JSON(http.StatusOK, statusReports)
}