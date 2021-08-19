package api

import (
	. "coffee_backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//获得所有记录
func GetProducts(c *gin.Context) {
	p := Product{}
	name := c.Request.FormValue("name")
	rs, _ := p.GetRows(name)
	// rs, _ := p.GetRows("", "", "")
	c.JSON(http.StatusOK, gin.H{
		"result": rs,
	})
}

func GetRecommendProducts(c *gin.Context) {
	p := Product{}
	rs, _ := p.GetRecommendRows()
	c.JSON(http.StatusOK, gin.H{
		"result": rs,
	})
}

//获得一条记录
func GetProduct(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	p := Product{
		Id: id,
	}
	rs, _ := p.GetRow()
	c.JSON(http.StatusOK, gin.H{
		"result": rs,
	})
}

//增加一条记录
func AddProduct(c *gin.Context) {
	var p Product
	c.BindJSON(&p)
	newSeq := GetSeq("PRODUCT")
	product := Product{
		Id:    newSeq,
		Name:  p.Name,
		Unit:  p.Unit,
		Price: p.Price,
		Image: p.Image,
		Memo:  p.Memo,
		State: p.State,
	}
	id := product.Create()
	res := "Error"
	if id == 0 {
		res = "Success"
	}
	c.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func UpdateProduct(c *gin.Context) {
	var p Product
	err := c.BindJSON(&p)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
	}
	product := Product{
		Id:    p.Id,
		Name:  p.Name,
		Unit:  p.Unit,
		Price: p.Price,
		Memo:  p.Memo,
		State: p.State,
		Image: p.Image,
	}
	id := product.Update()
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
func DelProduct(c *gin.Context) {
	ids := c.Request.FormValue("id")
	id, _ := strconv.Atoi(ids)
	row := Delete(id)
	msg := fmt.Sprintf("delete successful %d", row)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
