package serverinfo

import (
	"github.com/PittYao/stream_gateway/components/mysql"
	"gorm.io/gorm"
	"time"
)

type ServerInfo struct {
	gorm.Model
	Ip            string
	Port          string
	VideoType     string
	TaskIngNum    int
	FileSavePath  string
	DiskConnect   bool
	TotalCapacity string
	UsedCapacity  string
	FreeCapacity  string
	UsedPercent   string
	NginxRunning  bool
	HeartbeatTime time.Time
	Cpu           string
}

func GetAllServer() []*ServerInfo {
	var servers []*ServerInfo
	mysql.Instance.Find(&servers)
	return servers
}
