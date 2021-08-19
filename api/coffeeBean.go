package api

import (
	. "coffee_backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//获得所有记录
func GetCoffeeBeans(c *gin.Context) {
	cbs := CoffeeBean{}
	name := c.Request.FormValue("name")
	rs, _ := cbs.GetCoffeeBeans(name)
	c.JSON(http.StatusOK, gin.H{
		"result": rs,
	})
}

//获得一条记录
func GetCoffeeBean(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	cb := CoffeeBean{
		Id: id,
	}
	rs, _ := cb.GetCoffeeBean()
	c.JSON(http.StatusOK, gin.H{
		"result": rs,
	})
}

//增加一条记录
func AddCoffeeBean(c *gin.Context) {
	var coffeeBean CoffeeBean
	c.BindJSON(&coffeeBean)
	newSeq := GetSeq("COFFEEBEAN")
	cb := CoffeeBean{
		Id:           newSeq,
		Origin:       coffeeBean.Origin,
		Manor:        coffeeBean.Manor,
		Name:         coffeeBean.Name,
		Fermentation: coffeeBean.Fermentation,
		Roast:        coffeeBean.Roast,
		Acidity:      coffeeBean.Acidity,
		Flavor:       coffeeBean.Flavor,
		Tasting:      coffeeBean.Tasting,
		Description:  coffeeBean.Description,
		State:        coffeeBean.State,
		Store_Id:     coffeeBean.Store_Id,
		Image:        coffeeBean.Image,
		Buy_date:     coffeeBean.Buy_date,
		Weight:       coffeeBean.Weight,
		Price:        coffeeBean.Price,
	}
	id := cb.CreateCoffeeBean()
	res := "Error"
	if id == 0 {
		res = "Success"
	}
	c.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func UpdateCoffeeBean(c *gin.Context) {
	var coffeeBean CoffeeBean
	err := c.BindJSON(&coffeeBean)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
	}
	cb := CoffeeBean{
		Id:           coffeeBean.Id,
		Origin:       coffeeBean.Origin,
		Manor:        coffeeBean.Manor,
		Name:         coffeeBean.Name,
		Fermentation: coffeeBean.Fermentation,
		Roast:        coffeeBean.Roast,
		Acidity:      coffeeBean.Acidity,
		Flavor:       coffeeBean.Flavor,
		Tasting:      coffeeBean.Tasting,
		Description:  coffeeBean.Description,
		State:        coffeeBean.State,
		Store_Id:     coffeeBean.Store_Id,
		Image:        coffeeBean.Image,
		Buy_date:     coffeeBean.Buy_date,
		Weight:       coffeeBean.Weight,
		Price:        coffeeBean.Price,
	}
	id := cb.UpdateCoffeeBean()
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
func DelCoffeeBean(c *gin.Context) {
	ids := c.Request.FormValue("id")
	id, _ := strconv.Atoi(ids)
	row := DeleteCoffeeBean(id)
	msg := fmt.Sprintf("delete successful %d", row)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
