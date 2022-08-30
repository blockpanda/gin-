package itying

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItingController struct {
}

func (con ItingController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg": "我是一个msg",
	})
}
func (con ItingController) News(c *gin.Context) {
	c.String(200, "新闻")
}
