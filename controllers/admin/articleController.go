package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
}

func (con ArticleController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "编辑文章111111111111")
}
func (con ArticleController) Add(c *gin.Context) {
	c.String(http.StatusOK, "文章用户222222222")
}
func (con ArticleController) Dele(c *gin.Context) {
	c.String(http.StatusOK, "删除文章2223333222222")
}
