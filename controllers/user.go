package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	m "github.com/mgroast/ecmonster/models"
)

func GetUsers(c *gin.Context) {
	u := m.User{}
	users := u.GetUsers()
	c.HTML(http.StatusOK, "user.html", users)
}
