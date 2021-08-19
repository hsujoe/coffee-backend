package api

import (
	. "coffee_backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//获得所有记录
func GetRunShops(c *gin.Context) {
	runShop := RunShop{}
	name := c.Request.FormValue("name")
	rs, _ := runShop.GetRunShops(name)
	c.JSON(http.StatusOK, gin.H{
		"result": rs,
	})
}

//获得一条记录
func GetRunShop(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	runShop := RunShop{
		Id: id,
	}
	rs, _ := runShop.GetRunShop()
	c.JSON(http.StatusOK, gin.H{
		"result": rs,
	})
}

//增加一条记录
func AddRunShop(c *gin.Context) {
	var runShop RunShop
	c.BindJSON(&runShop)
	newSeq := GetSeq("RUNSHOP")
	rShop := RunShop{
		Id:         newSeq,
		Store_id:   runShop.Store_id,
		Product_id: runShop.Product_id,
		Tasting:    runShop.Tasting,
		Memo:       runShop.Memo,
		State:      runShop.State,
	}
	id := rShop.CreateRunShop()
	res := "Error"
	if id == 0 {
		res = "Success"
	}
	c.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func UpdateRunShop(c *gin.Context) {
	var runShop RunShop
	err := c.BindJSON(&runShop)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
	}
	rShop := RunShop{
		Id:         runShop.Id,
		Store_id:   runShop.Store_id,
		Product_id: runShop.Product_id,
		Tasting:    runShop.Tasting,
		Memo:       runShop.Memo,
		State:      runShop.State,
	}
	id := rShop.UpdateRunShop()
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
func DelRunShop(c *gin.Context) {
	ids := c.Request.FormValue("id")
	id, _ := strconv.Atoi(ids)
	row := DeleteRunShop(id)
	msg := fmt.Sprintf("delete successful %d", row)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
