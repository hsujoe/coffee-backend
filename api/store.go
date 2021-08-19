package api

import (
	. "coffee_backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//获得所有记录
func GetStores(c *gin.Context) {
	s := Store{}
	name := c.Request.FormValue("name")
	rs, _ := s.GetStores(name)
	c.JSON(http.StatusOK, gin.H{
		"result": rs,
	})
}

//获得所有记录
func GetRecommendStores(c *gin.Context) {
	s := Store{}
	rs, _ := s.GetRecommendStores()
	c.JSON(http.StatusOK, gin.H{
		"result": rs,
	})
}

//获得一条记录
func GetStore(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	s := Store{
		Id: id,
	}
	rs, _ := s.GetStore()
	c.JSON(http.StatusOK, gin.H{
		"result": rs,
	})
}

//增加一条记录
func AddStore(c *gin.Context) {
	var s Store
	c.BindJSON(&s)
	newSeq := GetSeq("STORE")
	store := Store{
		Id:      newSeq,
		Name:    s.Name,
		Tel:     s.Tel,
		Address: s.Address,
		Page:    s.Page,
		Fb:      s.Fb,
		Ig:      s.Ig,
		Image:   s.Image,
		Memo:    s.Memo,
		State:   s.State,
	}
	id := store.CreateStore()
	res := "Error"
	if id == 0 {
		res = "Success"
	}
	c.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func UpdateStore(c *gin.Context) {
	var s Store
	err := c.BindJSON(&s)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
	}
	store := Store{
		Id:      s.Id,
		Name:    s.Name,
		Tel:     s.Tel,
		Address: s.Address,
		Page:    s.Page,
		Fb:      s.Fb,
		Ig:      s.Ig,
		Memo:    s.Memo,
		State:   s.State,
		Image:   s.Image,
	}
	id := store.UpdateStore()
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
func DelStore(c *gin.Context) {
	ids := c.Request.FormValue("id")
	id, _ := strconv.Atoi(ids)
	row := DeleteStore(id)
	msg := fmt.Sprintf("delete successful %d", row)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
