package service

import (
	"fmt"
	"github.com/PittYao/stream_gateway/components/gin/response"
	"github.com/PittYao/stream_gateway/components/log"
	"github.com/PittYao/stream_gateway/helper"
	"github.com/PittYao/stream_gateway/internal/consts"
	"github.com/PittYao/stream_gateway/internal/dto"
	"github.com/PittYao/stream_gateway/internal/httpclient"
	"github.com/PittYao/stream_gateway/internal/model/ipserver"
	"github.com/PittYao/stream_gateway/internal/model/roommix3"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// StartMix3 godoc
// @Summary 开始
// @Tags 三合一
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Param startMix3Req body dto.StartMix3Req true " "
// @Router /api/v1/mix/transform/save/start [post]
func StartMix3(c *gin.Context) {
	var req dto.StartMix3Req
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, err.Error())
		return
	}

	// 查询rtspUrl的ip是否指定执行服务器
	rtspUrlMiddle := req.RtspUrlMiddle
	rtspUrlLeft := req.RtspUrlLeft
	rtspUrlRight := req.RtspUrlRight

	rtspUrls, err := helper.GetIpDirPathFormRtspUrls(",", rtspUrlMiddle, rtspUrlLeft, rtspUrlRight)
	if err != nil {
		log.L.Error(err.Error(), zap.Any("req", req))
		response.Err(c, err.Error())
		return
	}
	ipServer := ipserver.GetByCameraIpAndVideoType(rtspUrls, consts.Mix3)
	if ipServer == nil {
		log.L.Info("摄像头没有指定执行服务器ip", zap.Any("req", req))
		response.Err(c, "摄像头没有指定执行服务器ip")
		return
	}

	// 查询该rtsp任务是否已经在指定服务器上运行
	encodeRtspUrlMiddle := helper.EncodeRtspUrl(rtspUrlMiddle)
	encodeRtspUrlLeft := helper.EncodeRtspUrl(rtspUrlLeft)
	encodeRtspUrlRight := helper.EncodeRtspUrl(rtspUrlRight)

	serverHost := ipServer.ServerIp
	mix3s := roommix3.ListByIpAndRtspUrlsAndFfmpegSaveState(serverHost, encodeRtspUrlMiddle, encodeRtspUrlLeft, encodeRtspUrlRight, consts.RunIng)
	if len(mix3s) != 0 {
		log.L.Info("该rtsp已经在指定服务器上运行", zap.Any("req", req), zap.String("serverHost", serverHost))
		mix3 := mix3s[0]
		startRsp := &dto.StartRsp{
			TaskId:  mix3.ID,
			RtmpUrl: helper.GetRtmpUrlByIp(mix3.Ip, mix3.RtspUrlMiddle),
		}
		response.OKMsg(c, fmt.Sprintf("该rtsp已经在指定服务器:%s上运行", serverHost), startRsp)
		return
	}

	// 转发请求
	redirectUrl := helper.RedirectUrlBuilder(serverHost, consts.Mix3Port, fmt.Sprintf("/%s%s", consts.Mix3, consts.Start))
	log.L.Sugar().Infof("redirectUrl:%s", redirectUrl)
	c.Redirect(http.StatusPermanentRedirect, redirectUrl)
}

// StopMix3 godoc
// @Summary 停止
// @Tags 三合一
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Param stopReq body dto.StopReq true " "
// @Router /api/v1/mix/transform/save/stop [post]
func StopMix3(c *gin.Context) {
	var stopReq dto.StopReq
	if err := c.ShouldBindJSON(&stopReq); err != nil {
		response.Err(c, err.Error())
		return
	}

	// 查询该任务在哪个服务器执行
	mix3, err := roommix3.GetById(stopReq.TaskId)
	if err != nil {
		response.Err(c, err.Error())
		return
	}

	// 重定向到指定服务器
	serverHost := mix3.Ip
	redirectUrl := helper.RedirectUrlBuilder(serverHost, consts.Mix3Port, fmt.Sprintf("/%s%s", consts.Mix3, consts.Stop))
	c.Redirect(http.StatusPermanentRedirect, redirectUrl)
}

// StopAllMix3 godoc
// @Summary 停止所有
// @Tags 三合一
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/mix/transform/save/stopAll [post]
func StopAllMix3(c *gin.Context) {
	mix3IpServers := ipserver.ListByVideoType(consts.Mix3)

	if len(mix3IpServers) == 0 {
		response.Err(c, "没有配置三合一画面服务器ip")
		return
	}

	var stops []*dto.StopAllResp

	for _, server := range mix3IpServers {
		// 重定向
		stopAllUrl := helper.RedirectUrlBuilder(server.ServerIp, consts.Mix3Port, fmt.Sprintf("/%s%s", consts.Mix3, consts.StopAll))
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

	response.OKMsg(c, "关闭三合一画面任务成功", stops)
}

// RebootAllMix3 godoc
// @Summary 重启所有
// @Tags 三合一
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/mix/transform/save/reboot [post]
func RebootAllMix3(c *gin.Context) {
	// 查询所有的三合一画面服务
	mix3IpServers := ipserver.ListByVideoType(consts.Mix3)

	if len(mix3IpServers) == 0 {
		response.Err(c, "DB没有配置三合一画面服务器ip")
		return
	}

	var lists []*dto.ClientResponse

	for _, server := range mix3IpServers {
		// 转发请求
		serverHost := server.ServerIp
		rebootUrl := helper.RedirectUrlBuilder(serverHost, consts.Mix3Port, fmt.Sprintf("/%s%s", consts.Mix3, consts.RebootAll))
		err, resp := httpclient.RebootHttpClient(rebootUrl)

		list := &dto.ClientResponse{
			ServerIp: serverHost,
		}
		if err == nil {
			list.Data = resp.Data
		}
		lists = append(lists, list)
	}

	response.OKMsg(c, "重启所有三合一画面任务成功,各个服务器中异常的任务id如下", lists)
}

// ListAllMix3 godoc
// @Summary 查询所有
// @Tags 三合一
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/mix/transform/save/list [post]
func ListAllMix3(c *gin.Context) {
	// 查询所有的3合一画面服务
	mix3IpServers := ipserver.ListByVideoType(consts.Mix3)

	if len(mix3IpServers) == 0 {
		response.Err(c, "DB没有配置3合一画面服务器ip")
		return
	}

	var lists []*dto.ClientResponse

	for _, server := range mix3IpServers {
		// 转发请求
		serverHost := server.ServerIp
		listUrl := helper.RedirectUrlBuilder(serverHost, consts.Mix3Port, fmt.Sprintf("/%s%s", consts.Mix3, consts.GetAll))
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

	response.OKMsg(c, "查询所有3合一画面任务成功,各个服务器中异常的任务id如下", lists)
}
