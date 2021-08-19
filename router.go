package main

import (
	. "coffee_backend/api"
	"coffee_backend/middleware"

	"net/http"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	// http.HandleFunc("/", receiveClientRequest)
	router.Use(Cors())
	router.POST("/login", Login)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.NoCache())
	router.Use(middleware.Options())
	router.Use(middleware.Secure())
	// router.Use(mw...)
	router.Use(middleware.AuthJWT()) // 添加认证
	// router.GET("/", HomePage)        //http://localhost:8806

	//路由群组
	routers := router.Group("")
	{
		routers.GET("/user/:id", GetUser)

		routers.GET("/product", GetProducts)                          //http://localhost:8806/products
		routers.GET("/product/:id", GetProduct)                       //http://localhost:8806/products/{id}
		routers.GET("/products/recommend/data", GetRecommendProducts) //http://localhost:8806/products
		routers.PUT("/product/add", AddProduct)                       //http://localhost:8806/products/add
		routers.POST("/product/update", UpdateProduct)
		routers.POST("/product/del", DelProduct)

		routers.GET("/store", GetStores)
		routers.PUT("/store/add", AddStore)
		routers.GET("/stores/recommend/data", GetRecommendStores)
		routers.GET("/store/:id", GetStore)
		routers.POST("/store/update", UpdateStore)
		routers.POST("/store/del", DelStore)

		routers.GET("/coffeeBeans", GetCoffeeBeans)
		routers.PUT("/coffeeBean/add", AddCoffeeBean)
		routers.GET("/coffeeBean/:id", GetCoffeeBean)
		routers.POST("/coffeeBean/update", UpdateCoffeeBean)
		routers.POST("/coffeeBean/del", DelCoffeeBean)

		routers.GET("/selectItem/:dataName", GetStoreItem)

		routers.GET("/brew", GetBrews)
		routers.PUT("/brew/add", AddBrew)
		routers.GET("/brew/:id", GetBrew)
		routers.POST("/brew/update", UpdateBrew)
		routers.POST("/brew/del", DelBrew)

		routers.GET("/runShops", GetRunShops)
		routers.PUT("/runShop/add", AddRunShop)
		routers.GET("/runShop/:id", GetRunShop)
		routers.POST("/runShop/update", UpdateRunShop)
		routers.POST("/runShop/del", DelRunShop)
	}

	return router
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// func receiveClientRequest(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
// 	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
// 	w.Header().Set("content-type", "application/json")             //返回数据格式是json

// 	r.ParseForm()
// 	fmt.Println("收到客户端请求: ", r.Form)
// }
