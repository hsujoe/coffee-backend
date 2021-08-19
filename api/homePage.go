package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//index
func HomePage(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}
