package api

import (
	. "coffee_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u LoginToken
	c.BindJSON(&u)
	// username := c.Request.FormValue("username")
	// password := c.Request.FormValue("password")
	// if err == nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": err,
	// 	})
	// }
	//checkUser
	re, _ := CheckUser(u.UserId, u.Password)
	// re, _ := CheckUser(username, password)
	if re.ID == 0 {
		// err := errors.New("帳密錯誤")
		c.JSON(http.StatusOK, gin.H{
			"result": re,
		})
		c.Abort()
		return
	}

	// 確認無誤签发 token
	t, _ := Sign(u.UserId, u.Password)
	// t, _ := Sign(username, password)
	c.JSON(http.StatusOK, gin.H{
		"data": re, "token": t,
	})
}

func GetUser(c *gin.Context) {
	// ids := c.Param("id")
	// id, _ := strconv.Atoi(ids)
	LU := LoginToken{
		ID: 1,
	}
	//checkUser
	re, _ := LU.GetUser()

	if re.ID == 0 {
		// err := errors.New("帳密錯誤")
		c.JSON(http.StatusOK, gin.H{
			"result": re,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": re,
	})
}
