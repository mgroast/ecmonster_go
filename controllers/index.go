package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	m "github.com/mgroast/ecmonster/models"
)

func GetIndex(c *gin.Context) {
	p := m.Product{}
	products := p.FindNewProducts()	
	c.HTML(http.StatusOK, "index.html", gin.H{"products": products})
}
