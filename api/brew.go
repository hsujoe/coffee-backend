package api

import (
	. "coffee_backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//获得所有记录
func GetBrews(c *gin.Context) {
	cbs := Brew{}
	name := c.Request.FormValue("name")
	rs, _ := cbs.GetBrews(name)
	c.JSON(http.StatusOK, gin.H{
		"result": rs,
	})
}

//获得一条记录
func GetBrew(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	cb := Brew{
		Id: id,
	}
	rs, _ := cb.GetBrew()
	c.JSON(http.StatusOK, gin.H{
		"result": rs,
	})
}

//增加一条记录
func AddBrew(c *gin.Context) {
	var brew Brew
	c.BindJSON(&brew)
	newSeq := GetSeq("BREW")
	b := Brew{
		Id:          newSeq,
		Grinder:     brew.Grinder,
		Scale:       brew.Scale,
		Steaming:    brew.Steaming,
		WaterCuts:   brew.WaterCuts,
		Totalwater:  brew.Totalwater,
		Description: brew.Description,
		Bean_Id:     brew.Bean_Id,
	}
	id := b.CreateBrew()
	res := "Error"
	if id == 0 {
		res = "Success"
	}
	c.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func UpdateBrew(c *gin.Context) {
	var brew Brew
	err := c.BindJSON(&brew)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
	}
	b := Brew{
		Id:          brew.Id,
		Grinder:     brew.Grinder,
		Scale:       brew.Scale,
		Steaming:    brew.Steaming,
		WaterCuts:   brew.WaterCuts,
		Totalwater:  brew.Totalwater,
		Description: brew.Description,
		Bean_Id:     brew.Bean_Id,
	}
	id := b.UpdateBrew()
	res := "Error"
	if id == 0 {
		res = "Success"
	}
	// resData := fmt.Sprintf("updated successful,%d", strconv.FormatBool(res))
	c.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

//删除一条记录
func DelBrew(c *gin.Context) {
	ids := c.Request.FormValue("id")
	id, _ := strconv.Atoi(ids)
	row := DeleteBrew(id)
	msg := fmt.Sprintf("delete successful %d", row)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
