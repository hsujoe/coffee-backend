package api

import (
	. "coffee_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//获得所有记录
func GetStoreItem(c *gin.Context) {
	dataName := c.Param("dataName")
	rs, _ := GetSelectItem(dataName)
	c.JSON(http.StatusOK, gin.H{
		"result": rs,
	})
}
