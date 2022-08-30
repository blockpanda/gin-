package routers

import (
	"gin/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", admin.IndexController{}.Index)
		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/edit", admin.ArticleController{}.Edit)
		adminRouters.GET("/article/dele", admin.ArticleController{}.Dele)

	}
}
