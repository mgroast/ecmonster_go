package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
