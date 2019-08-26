package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProduct(c *gin.Context) {
	c.HTML(http.StatusOK, "product.html", gin.H{})
}
