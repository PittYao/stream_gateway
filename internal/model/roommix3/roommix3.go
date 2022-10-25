package roommix3

import (
	"errors"
	"github.com/PittYao/stream_gateway/components/log"
	"github.com/PittYao/stream_gateway/components/mysql"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type RoomMix3 struct {
	gorm.Model
	RtspUrlMiddle            string
	RtspUrlLeft              string
	RtspUrlRight             string
	Temperature              string
	RoomName                 string
	Ip                       string
	Port                     string
	SavePath                 string
	FileRecentTime           *time.Time
	FfmpegTransformState     int
	FfmpegTransformCmd       string
	FfmpegTransformErrorMsg  string
	FfmpegTransformStartTime *time.Time
	FfmpegTransformCloseTime *time.Time
	FfmpegSaveState          int
	FfmpegSaveCmd            string
	FfmpegSaveErrorMsg       string
	FfmpegSaveStartTime      *time.Time
	FfmpegSaveCloseTime      *time.Time
	FfmpegStateLog           string
	TsFile                   string
	RebootRootID             uint
	RebootParentID           uint
	DisuseAt                 *time.Time
	M3u8Url                  string
}

// --- orm --- //

// Add 插入任务
func (r *RoomMix3) Add() error {
	create := mysql.Instance.Create(r)
	if create.Error != nil {
		log.L.Error("RoomMix3 新增mix3转流任务失败", zap.Error(create.Error))
		return errors.New(" 新增mix3转流任务失败")
	}
	//新增mix3转流任务失败   {"error": "Error 1292: Incorrect datetime value: '0000-00-00' for column 'file_recent_time' at row 1"}
	return nil
}

// Update 更新
func (r *RoomMix3) Update() error {
	save := mysql.Instance.Save(&r)
	if save.Error != nil {
		log.L.Error("RoomMix3 更新失败", zap.Error(save.Error))
		return errors.New("RoomMix3 更新失败")
	}
	return save.Error

}

// Delete
func (r *RoomMix3) Delete() error {
	save := mysql.Instance.Delete(&r)
	if save.Error != nil {
		log.L.Error("RoomMix3 删除失败", zap.Error(save.Error))
		return errors.New("RoomMix3 删除失败")
	}
	return save.Error

}

// GetById id查询
func GetById(id uint) (*RoomMix3, error) {
	var roomMix3 RoomMix3
	mysql.Instance.First(&roomMix3, id)

	if roomMix3.ID == 0 {
		log.L.Sugar().Errorf("没有查询到该3合一画面任务 id:%d", id)
		return nil, errors.New("没有查询到该3合一画面任务")
	}

	return &roomMix3, nil
}

func ListByIpAndRtspUrlsAndFfmpegSaveState(ip, rtspUrlMiddle, rtspUrlLeft, rtspUrlRight string, ffmpegSaveState int) (roomMix3s []RoomMix3) {
	mysql.Instance.Where("ip = ? and rtsp_url_middle = ? and rtsp_url_left = ? and rtsp_url_right = ? and ffmpeg_save_state = ?",
		ip, rtspUrlMiddle, rtspUrlLeft, rtspUrlRight, ffmpegSaveState).Find(&roomMix3s)
	return
}
