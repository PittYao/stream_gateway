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
	"github.com/PittYao/stream_gateway/internal/model/roommix4"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// StartMix4 godoc
// @Summary 开始
// @Tags 四合一
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Param req body dto.StartMix4Req true " "
// @Router /api/v1/mix/transform/save/41/start [post]
func StartMix4(c *gin.Context) {
	var req dto.StartMix4Req
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Err(c, err.Error())
		return
	}

	// 查询rtspUrl的ip是否指定执行服务器
	rtspUrlMiddle := req.RtspUrlMiddle
	rtspUrlSmallOne := req.RtspUrlSmallOne
	rtspUrlSmallTwo := req.RtspUrlSmallTwo
	rtspUrlSmallThree := req.RtspUrlSmallThree

	rtspUrls, err := helper.GetIpDirPathFormRtspUrls(",", rtspUrlMiddle, rtspUrlSmallOne, rtspUrlSmallTwo, rtspUrlSmallThree)
	if err != nil {
		log.L.Error(err.Error(), zap.Any("req", req))
		response.Err(c, err.Error())
		return
	}
	ipServer := ipserver.GetByCameraIpAndVideoType(rtspUrls, consts.Mix4)
	if ipServer == nil {
		log.L.Info("摄像头没有指定执行服务器ip", zap.Any("req", req))
		response.Err(c, "摄像头没有指定执行服务器ip")
		return
	}

	// 查询该rtsp任务是否已经在指定服务器上运行
	encodeRtspUrlMiddle := helper.EncodeRtspUrl(rtspUrlMiddle)
	encodeRtspUrlSmallOne := helper.EncodeRtspUrl(rtspUrlSmallOne)
	encodeRtspUrlSmallTwo := helper.EncodeRtspUrl(rtspUrlSmallTwo)
	encodeRtspUrlSmallThree := helper.EncodeRtspUrl(rtspUrlSmallThree)

	serverHost := ipServer.ServerIp
	mix4s := roommix4.ListByIpAndRtspUrlsAndFfmpegSaveState(serverHost, encodeRtspUrlMiddle, encodeRtspUrlSmallOne, encodeRtspUrlSmallTwo, encodeRtspUrlSmallThree, consts.RunIng)
	if len(mix4s) != 0 {
		log.L.Info("该rtsp已经在指定服务器上运行", zap.Any("req", req), zap.String("serverHost", serverHost))
		startRsp := &dto.StartRsp{
			TaskId:  mix4s[0].ID,
			RtmpUrl: "",
		}
		response.OKMsg(c, fmt.Sprintf("该rtsp已经在指定服务器:%s上运行", serverHost), startRsp)
		return
	}

	// 转发请求
	redirectUrl := helper.RedirectUrlBuilder(serverHost, consts.Mix4Port, fmt.Sprintf("/%s%s", consts.Mix4, consts.Start))
	c.Redirect(http.StatusPermanentRedirect, redirectUrl)
}

// StopMix4 godoc
// @Summary 停止
// @Tags 四合一
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Param stopReq body dto.StopReq true " "
// @Router /api/v1/mix/transform/save/41/stop [post]
func StopMix4(c *gin.Context) {
	var stopReq dto.StopReq
	if err := c.ShouldBindJSON(&stopReq); err != nil {
		response.Err(c, err.Error())
		return
	}

	// 查询该任务在哪个服务器执行
	mix4, err := roommix4.GetById(stopReq.TaskId)
	if err != nil {
		response.Err(c, err.Error())
		return
	}

	// 重定向到指定服务器
	serverHost := mix4.Ip
	redirectUrl := helper.RedirectUrlBuilder(serverHost, consts.Mix4Port, fmt.Sprintf("/%s%s", consts.Mix4, consts.Stop))
	c.Redirect(http.StatusPermanentRedirect, redirectUrl)
}

// StopAllMix4 godoc
// @Summary 停止所有
// @Tags 四合一
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/mix/transform/save/41/stopAll [post]
func StopAllMix4(c *gin.Context) {
	mix4IpServers := ipserver.ListByVideoType(consts.Mix4)

	if len(mix4IpServers) == 0 {
		response.Err(c, "没有配置四合一画面服务器ip")
		return
	}

	var stops []*dto.StopAllResp

	for _, server := range mix4IpServers {
		// 重定向
		stopAllUrl := helper.RedirectUrlBuilder(server.ServerIp, consts.Mix4Port, fmt.Sprintf("/%s%s", consts.Mix4, consts.StopAll))
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

	response.OKMsg(c, "关闭四合一画面任务成功", stops)
}

// RebootAllMix4 godoc
// @Summary 重启所有
// @Tags 四合一
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/mix/transform/save/41/reboot [post]
func RebootAllMix4(c *gin.Context) {
	// 查询所有的四合一画面服务
	mix4IpServers := ipserver.ListByVideoType(consts.Mix4)

	if len(mix4IpServers) == 0 {
		response.Err(c, "DB没有配置四合一画面服务器ip")
		return
	}

	var lists []*dto.ClientResponse

	for _, server := range mix4IpServers {
		// 转发请求
		serverHost := server.ServerIp
		rebootUrl := helper.RedirectUrlBuilder(serverHost, consts.Mix4Port, fmt.Sprintf("/%s%s", consts.Mix4, consts.RebootAll))
		err, resp := httpclient.RebootHttpClient(rebootUrl)

		list := &dto.ClientResponse{
			ServerIp: serverHost,
		}
		if err == nil {
			list.Data = resp.Data
		}
		lists = append(lists, list)
	}

	response.OKMsg(c, "重启所有四合一画面任务成功,各个服务器中异常的任务id如下", lists)
}

// ListAllMix4 godoc
// @Summary 查询所有
// @Tags 四合一
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/mix/transform/save/41/list [post]
func ListAllMix4(c *gin.Context) {
	// 查询所有的3合一画面服务
	mix4IpServers := ipserver.ListByVideoType(consts.Mix4)

	if len(mix4IpServers) == 0 {
		response.Err(c, "DB没有配置3合一画面服务器ip")
		return
	}

	var lists []*dto.ClientResponse

	for _, server := range mix4IpServers {
		// 转发请求
		serverHost := server.ServerIp
		listUrl := helper.RedirectUrlBuilder(serverHost, consts.Mix4Port, fmt.Sprintf("/%s%s", consts.Mix4, consts.GetAll))
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

	response.OKMsg(c, "查询所有四合一画面任务成功,各个服务器中异常的任务id如下", lists)
}
