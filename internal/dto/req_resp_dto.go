package dto

// StartMix4Req 四合一请求体
type StartMix4Req struct {
	RtspUrlMiddle     string `json:"rtspUrlMiddle" binding:"required" example:"rtsp://admin:CEBON123@192.168.99.115"`
	RtspUrlSmallOne   string `json:"rtspUrlSmallOne" binding:"required" example:"rtsp://admin:cebon61332433@192.168.99.112"`
	RtspUrlSmallTwo   string `json:"rtspUrlSmallTwo" binding:"required" example:"rtsp://admin:cebon61332433@192.168.99.215"`
	RtspUrlSmallThree string `json:"rtspUrlSmallThree" binding:"required" example:"rtsp://admin:cebon61332433@192.168.99.215"`
	Temperature       string `json:"temperature" example:""`
	RoomName          string `json:"roomName" example:"1"`
}

// StartMix3Req 三合一请求体
type StartMix3Req struct {
	RtspUrlMiddle string `json:"rtspUrlMiddle" binding:"required" example:"rtsp://admin:cebon61332433@192.168.99.215:554/cam/realmonitor?channel=1&subtype=0"`
	RtspUrlLeft   string `json:"rtspUrlLeft" binding:"required" example:"rtsp://admin:cebon61332433@192.168.99.215:554/cam/realmonitor?channel=1&subtype=1"`
	RtspUrlRight  string `json:"rtspUrlRight" binding:"required" example:"rtsp://admin:cebon61332433@192.168.99.215:554/cam/realmonitor?channel=1&subtype=1"`
	Temperature   string `json:"temperature" example:""`
	RoomName      string `json:"roomName" example:"1"`
}

// StartReq 单个摄像头请求体
type StartReq struct {
	RtspUrl string `json:"rtspUrl" binding:"required" example:"rtsp://admin:cebon61332433@192.168.99.215:554/cam/realmonitor?channel=1&subtype=0"`
}

// StartRsp 开始存储响应体
type StartRsp struct {
	TaskId  uint   `json:"taskId"`
	RtmpUrl string `json:"rtmpUrl"`
}

// StopReq 结束存储响应体
type StopReq struct {
	RtmpUrl string `json:"rtmpUrl" binding:"required"`
	TaskId  uint   `json:"taskId"  binding:"required"`
}

// CopyReq 拷贝临时m3u8地址
type CopyReq struct {
	Id uint `json:"id" example:"1" binding:"required"`
}

// GetAllRsp 获取正在进行的任务
type GetAllRsp struct {
	RtspUrl  string `json:"rtspUrl"`
	RoomName string `json:"roomName"`
}

// StopAllResp 停止所有响应体
type StopAllResp struct {
	ServerIp string `json:"serverIp"`
	Msg      string `json:"msg"`
}

type ClientResponse struct {
	ServerIp string      `json:"serverIp"`
	Data     interface{} `json:"data"`
}

type UpgradeReq struct {
	FileUrl string `json:"fileUrl"` // 升级文件下载地址
}
