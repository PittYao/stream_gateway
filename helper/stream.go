package helper

import (
	"fmt"
	"github.com/PittYao/stream_push_save/components/config"
	"github.com/PittYao/stream_push_save/components/log"
	"github.com/PittYao/stream_push_save/internal/consts"
	"github.com/duke-git/lancet/datetime"
	"strings"
	"time"
)

// EncodeRtspUrl 替换rtspUrl中的&为 %26
func EncodeRtspUrl(rtspUrl string) string {
	index := strings.Index(rtspUrl, "&")
	if index != -1 {
		rtspUrl = strings.Replace(rtspUrl, "&", "%26", -1)
	}
	return rtspUrl
}

// CmdString2Array cmd命令由string转为string数组
func CmdString2Array(strCmd string) []string {
	args := strings.Split(strCmd, " ")
	var resultArgs []string
	for _, v := range args {
		if v != "" {
			resultArgs = append(resultArgs, v)
		}
	}
	return resultArgs
}

// GetRebootTime 计算重启时间，用于定时重启
func GetRebootTime() time.Time {
	//minute := config.C.Video.M3u8MaxTime * 60 // 单位分钟
	return datetime.AddMinute(time.Now(), 1)
}

// GetRtmpUrl 由rtsp地址获取rtmp地址
func GetRtmpUrl(rtspUrl string) string {
	rtmpUrl := "rtmp://" + config.C.Server.Ip + ":1935/live/" + Md5ByString(rtspUrl)
	return rtmpUrl
}

// GetM3u8Url 获取m3u8网络地址
func GetM3u8Url(ip, filePath string) string {
	filePath = strings.ReplaceAll(filePath, "\\", "/")

	index := strings.Index(filePath, "/")
	if index == -1 {
		log.L.Sugar().Errorf("m3u8Url生成异常 filePath:%s", filePath)
		return ""
	}

	m3u8Url := fmt.Sprintf("http://%s%s%s/%s", ip, consts.M3u8UrlPort, filePath[index:], consts.M3u8File)

	// http://192.168.99.19:8880/videodata/publicSingle/192.168.99.117/2022.04.21-13.08.08/playlist.m3u8
	return m3u8Url
}

// GetTempM3u8Url 获取m3u8临时网络地址
func GetTempM3u8Url(ip, filePath, m3u8FileName string) string {
	filePath = strings.ReplaceAll(filePath, "\\", "/")

	index := strings.Index(filePath, "/")
	if index == -1 {
		log.L.Sugar().Errorf("m3u8Url生成异常 filePath:%s", filePath)
		return ""
	}

	m3u8Url := fmt.Sprintf("http://%s%s%s/%s", ip, consts.M3u8UrlPort, filePath[index:], m3u8FileName)

	// http://192.168.99.19:8880/videodata/publicSingle/192.168.99.117/2022.04.21-13.08.08/playlist.m3u8
	return m3u8Url
}
