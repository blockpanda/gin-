package routers

import (
	itying "gin/controllers/iting"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/") //路由分组
	{
		defaultRouters.GET("/", itying.ItingController{}.Index)
		defaultRouters.GET("/news", itying.ItingController{}.News)
	}

}
