package router

import (
	"github.com/gin-gonic/gin"
	"cdn-server/controller"
)

func ConfigRouter(r gin.IRouter) {
	//app 接口
	app := r.Group("/api/")
	{
		app.GET("libraries", controller.LibrariesList)
		app.GET("libraries/:name", controller.LibrariesInfo)
	}

}
