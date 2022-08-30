package routers

import (
	apiCon "gin/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", apiCon.ApiController{}.Index)

		apiRouters.GET("/userlist", apiCon.ApiController{}.Apilist)
	}

}
