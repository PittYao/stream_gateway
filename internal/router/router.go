// 在这个文件中注册 URL handler

package router

import (
	"github.com/PittYao/stream_gateway/internal/service"
	"github.com/gin-gonic/gin"
)

// Routes 注册 API URL 路由
func Routes(app *gin.Engine) {
	recordGroup := app.Group("/api")
	{
		v1 := recordGroup.Group("/v1")
		{
			mix3Group := v1.Group("/mix/transform/save")
			{
				mix3Group.POST("/start", service.StartMix3)
				mix3Group.POST("/stop", service.StopMix3)
				mix3Group.POST("/stopAll", service.StopAllMix3)
				mix3Group.POST("/rebootAll", service.RebootAllMix3)
				mix3Group.POST("/listAll", service.ListAllMix3)
			}
		}

	}
}
