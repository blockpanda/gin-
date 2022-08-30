package apiCon

import "github.com/gin-gonic/gin"

type ApiController struct {
}

func (con ApiController) Index(c *gin.Context) {
	c.String(200, "API接口55555555555")
}
func (con ApiController) Apilist(c *gin.Context) {
	c.String(200, "userlistAPI接口555555555555555")
}
