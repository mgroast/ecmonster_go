package main

import(
	"github.com/gin-gonic/gin"
	c "github.com/mgroast/ecmonster/controllers"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/", c.GetIndex)
	r.Run(":8080")
}
