package ipserver

import (
	"github.com/PittYao/stream_gateway/components/mysql"
	"gorm.io/gorm"
	"time"
)

type IpServer struct {
	gorm.Model
	CameraIp      string
	ServerIp      string
	ServerPort    string
	FileSizeCheck bool
	VideoType     string
	PingTime      time.Time
	NetworkInfo   string
	DontSave      bool // 不存储
}

func GetByCameraIpAndVideoType(cameraIp, videoType string) (ipServer *IpServer) {
	mysql.Instance.Last(&ipServer, "camera_ip = ? and video_type = ? ", cameraIp, videoType)
	if ipServer.ID == 0 {
		return nil
	}
	return
}

func ListByVideoType(videoType string) (ipServers []IpServer) {
	mysql.Instance.Distinct("server_ip").Where("video_type = ?", videoType).Find(&ipServers)
	return
}
