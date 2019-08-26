package main

import(
	"github.com/gin-gonic/gin"
	c "github.com/mgroast/ecmonster/controllers"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("views/*")
	r.GET("/", c.GetIndex)
	r.GET("/login", c.GetLogin)
	r.GET("/product", c.GetProduct)
	r.GET("/user", c.GetUsers)
	r.Run(":8080")
}
