package service

import (
	"fmt"
	"github.com/PittYao/stream_gateway/components/gin/response"
	"github.com/PittYao/stream_gateway/components/log"
	"github.com/PittYao/stream_gateway/helper"
	"github.com/PittYao/stream_gateway/internal/consts"
	"github.com/PittYao/stream_gateway/internal/dto"
	"github.com/PittYao/stream_gateway/internal/httpclient"
	"github.com/PittYao/stream_gateway/internal/model/serverinfo"
	"github.com/gin-gonic/gin"
)

// Upgrade godoc
// @Summary 更新程序
// @Tags 更新程序
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Param req body dto.UpgradeReq true " "
// @Router  /api/v1/upgrade [post]
func Upgrade(c *gin.Context) {
	var req dto.UpgradeReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.Err(c, err.Error())
		return
	}

	servers := serverinfo.GetAllServer()
	if len(servers) == 0 {
		response.Err(c, "服务器列表为空,不升级")
		return
	}

	for _, server := range servers {
		var port string
		switch server.VideoType {
		case consts.Mix3:
			port = consts.Mix3Port
		case consts.Mix4:
			port = consts.Mix4Port
		case consts.Single:
			port = consts.SinglePort
		case consts.PublicSingle:
			port = consts.PublicPort
		}

		url := helper.RedirectUrlBuilder(server.Ip, port, fmt.Sprintf("%s", consts.Upgrade))
		err = httpclient.UpgradeHttpClient(url, req)
		if err != nil {
			log.L.Sugar().Errorf("下载程序包失败,url:%s", url)
		} else {
			log.L.Sugar().Infof("下载程序包成功,url:%s", url)
		}

	}

}
