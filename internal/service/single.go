package service

import (
	"fmt"
	"github.com/PittYao/stream_gateway/components/config"
	"github.com/PittYao/stream_gateway/components/gin/response"
	"github.com/PittYao/stream_gateway/components/log"
	"github.com/PittYao/stream_gateway/helper"
	"github.com/PittYao/stream_gateway/internal/consts"
	"github.com/PittYao/stream_gateway/internal/dto"
	"github.com/PittYao/stream_gateway/internal/httpclient"
	"github.com/PittYao/stream_gateway/internal/model/ipserver"
	"github.com/PittYao/stream_gateway/internal/model/roomrecordone"
	"github.com/duke-git/lancet/random"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// StartSingle godoc
// @Summary 开始
// @Tags 房间单画面
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Param req body dto.StartReq true " "
// @Router /api/v1/single/transform/save/start [post]
func StartSingle(c *gin.Context) {
	var req dto.StartReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, err.Error())
		return
	}

	// 从rtspUrl中获取ip
	cameraIp, err := helper.GetIpDirPathFormRtspUrls(",", req.RtspUrl)
	if err != nil {
		log.L.Error(err.Error(), zap.Any("req", req))
		response.Err(c, err.Error())
		return
	}
	// 查询rtspUrl的ip是否指定执行服务器
	ipServer := ipserver.GetByCameraIpAndVideoType(cameraIp, consts.Single)
	if ipServer == nil {
		log.L.Info("摄像头没有指定执行服务器ip", zap.Any("req", req))
		response.Err(c, "摄像头没有指定执行服务器ip")
		return
	}
	if ipServer.DontSave {
		startRsp := &dto.StartRsp{
			TaskId:  uint(random.RandInt(1, 20000)),
			RtmpUrl: helper.GetRtmpUrlByIp(config.C.Server.Ip, req.RtspUrl),
		}
		response.OKMsg(c, fmt.Sprintf("该地址配置不存储,返回随机值"), startRsp)
		return
	}

	// 查询该rtsp任务是否已经在指定服务器上运行
	encodeRtspUrl := helper.EncodeRtspUrl(req.RtspUrl)

	serverHost := ipServer.ServerIp
	singles := roomrecordone.ListByIpAndRtspUrlsAndFfmpegSaveState(serverHost, encodeRtspUrl, consts.RunIng)
	if len(singles) != 0 {
		log.L.Info("该rtsp已经在指定服务器上运行", zap.Any("req", req), zap.String("serverHost", serverHost))
		one := singles[0]
		startRsp := &dto.StartRsp{
			TaskId:  one.ID,
			RtmpUrl: helper.GetRtmpUrlByIp(one.Ip, one.RtspUrl),
		}
		response.OKMsg(c, fmt.Sprintf("该rtsp已经在指定服务器:%s上运行", serverHost), startRsp)
		return
	}

	// 转发请求
	redirectUrl := helper.RedirectUrlBuilder(serverHost, consts.SinglePort, fmt.Sprintf("/%s%s", consts.Single, consts.Start))
	c.Redirect(http.StatusPermanentRedirect, redirectUrl)
}

// StopSingle godoc
// @Summary 停止
// @Tags 房间单画面
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Param stopReq body dto.StopReq true " "
// @Router /api/v1/single/transform/save/stop [post]
func StopSingle(c *gin.Context) {
	var stopReq dto.StopReq
	if err := c.ShouldBindJSON(&stopReq); err != nil {
		response.Err(c, err.Error())
		return
	}

	if strings.Contains(stopReq.RtmpUrl, config.C.Server.Ip) {
		response.OKMsg(c, "该地址配置为不存储", nil)
		return
	}

	// 查询该任务在哪个服务器执行
	one, err := roomrecordone.GetById(stopReq.TaskId)
	if err != nil {
		response.Err(c, err.Error())
		return
	}

	// 重定向到指定服务器
	serverHost := one.Ip
	redirectUrl := helper.RedirectUrlBuilder(serverHost, consts.SinglePort, fmt.Sprintf("/%s%s", consts.Single, consts.Stop))
	c.Redirect(http.StatusPermanentRedirect, redirectUrl)
}

// StopAllSingle godoc
// @Summary 停止所有
// @Tags 房间单画面
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/single/transform/save/stopAll [post]
func StopAllSingle(c *gin.Context) {
	singleIpServers := ipserver.ListByVideoType(consts.Single)

	if len(singleIpServers) == 0 {
		response.Err(c, "没有配置房间画面服务器ip")
		return
	}

	var stops []*dto.StopAllResp

	for _, server := range singleIpServers {
		// 重定向
		stopAllUrl := helper.RedirectUrlBuilder(server.ServerIp, consts.SinglePort, fmt.Sprintf("/%s%s", consts.Single, consts.StopAll))
		err := httpclient.StopHttpClient(stopAllUrl)

		stop := &dto.StopAllResp{
			ServerIp: server.ServerIp,
		}
		if err != nil {
			stop.Msg = err.Error()
		} else {
			stop.Msg = "结束成功"
		}
		stops = append(stops, stop)
	}

	response.OKMsg(c, "关闭房间画面任务成功", stops)
}

// RebootAllSingle godoc
// @Summary 重启所有
// @Tags 房间单画面
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/single/transform/save/reboot [post]
func RebootAllSingle(c *gin.Context) {
	// 查询所有的房间画面服务
	singleIpServers := ipserver.ListByVideoType(consts.Single)

	if len(singleIpServers) == 0 {
		response.Err(c, "DB没有配置房间画面服务器ip")
		return
	}

	var lists []*dto.ClientResponse

	for _, server := range singleIpServers {
		// 转发请求
		serverHost := server.ServerIp
		rebootUrl := helper.RedirectUrlBuilder(serverHost, consts.SinglePort, fmt.Sprintf("/%s%s", consts.Single, consts.RebootAll))
		err, resp := httpclient.RebootHttpClient(rebootUrl)

		list := &dto.ClientResponse{
			ServerIp: serverHost,
		}
		if err == nil {
			list.Data = resp.Data
		}
		lists = append(lists, list)
	}

	response.OKMsg(c, "重启所有房间画面任务成功,各个服务器中异常的任务id如下", lists)
}

// ListAllSingle godoc
// @Summary 查询所有
// @Tags 房间单画面
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/single/transform/save/list [post]
func ListAllSingle(c *gin.Context) {
	SingleIpServers := ipserver.ListByVideoType(consts.Single)

	if len(SingleIpServers) == 0 {
		response.Err(c, "DB没有配置房间画面服务器ip")
		return
	}

	var lists []*dto.ClientResponse

	for _, server := range SingleIpServers {
		// 转发请求
		serverHost := server.ServerIp
		listUrl := helper.RedirectUrlBuilder(serverHost, consts.SinglePort, fmt.Sprintf("/%s%s", consts.Single, consts.GetAll))
		err, resp := httpclient.ListStreamHttpClient(listUrl)
		if err != nil {
			log.L.Sugar().Errorf("查询服务器:%s,异常", serverHost)
			continue
		}
		list := &dto.ClientResponse{
			ServerIp: serverHost,
		}
		if err == nil {
			list.Data = resp.Data
		}
		lists = append(lists, list)
	}

	response.OKMsg(c, "查询所有房间画面任务成功,各个服务器中异常的任务id如下", lists)
}
