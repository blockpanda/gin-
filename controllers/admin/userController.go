package admin

import "github.com/gin-gonic/gin"

type UserController struct {
	BaseContriller
}

func (con UserController) Index(c *gin.Context) {
	c.String(200, "用户列表111111111111")
	con.success(c)
}
func (con UserController) Add(c *gin.Context) {
	c.String(200, "增加用户222222222")
}
func (con UserController) Edit(c *gin.Context) {
	c.String(200, "编辑用户2223333222222")
}
