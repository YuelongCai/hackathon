package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BadRequest return 400 and error message
func BadRequest(c *gin.Context, errMsg string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"error": errMsg})
}

// InternalServerErrorRequest return 500 and error message
func InternalServerErrorRequest(c *gin.Context, errMsg string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"error": errMsg})
}

// SuccessRequest return 200 and interface
func SuccessRequest(c *gin.Context, body interface{}) {
	c.JSON(http.StatusOK, body)
}
