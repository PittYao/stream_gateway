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
				mix3Group.POST("/list", service.ListAllMix3)
				mix3Group.POST("/reboot", service.RebootAllMix3)
			}
			mix4Group := v1.Group("/mix/transform/save/41")
			{
				mix4Group.POST("/start", service.StartMix4)
				mix4Group.POST("/stop", service.StopMix4)
				mix4Group.POST("/stopAll", service.StopAllMix4)
				mix4Group.POST("/list", service.ListAllMix4)
				mix4Group.POST("/reboot", service.RebootAllMix4)
			}
			singleGroup := v1.Group("/single/transform/save")
			{
				singleGroup.POST("/start", service.StartSingle)
				singleGroup.POST("/stop", service.StopSingle)
				singleGroup.POST("/stopAll", service.StopAllSingle)
				singleGroup.POST("/list", service.ListAllSingle)
				singleGroup.POST("/reboot", service.RebootAllSingle)
			}
			publicSingleGroup := v1.Group("/other/transform/save")
			{
				publicSingleGroup.POST("/start", service.StartPublic)
				publicSingleGroup.POST("/stop", service.StopPublic)
				publicSingleGroup.POST("/stopAll", service.StopAllPublic)
				publicSingleGroup.POST("/list", service.ListAllPublic)
				publicSingleGroup.POST("/reboot", service.RebootAllPublic)
			}

			v1.POST("/upgrade", service.Upgrade)
		}

	}
}
