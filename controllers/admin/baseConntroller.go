package admin

import "github.com/gin-gonic/gin"

type BaseContriller struct {
}

func (con BaseContriller) success(c *gin.Context) {
	c.String(200, "成功")
}
func (con BaseContriller) error(c *gin.Context) {
	c.String(200, "失败")
}
