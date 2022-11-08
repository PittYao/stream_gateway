package roommix4

import (
	"errors"
	"github.com/PittYao/stream_gateway/components/log"
	"github.com/PittYao/stream_gateway/components/mysql"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type RoomMix4 struct {
	gorm.Model
	RtspUrlMiddle            string
	RtspUrlSmallOne          string
	RtspUrlSmallTwo          string
	RtspUrlSmallThree        string
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

// Add 插入任务
func (r *RoomMix4) Add() error {
	create := mysql.Instance.Create(r)
	if create.Error != nil {
		log.L.Error("新增mix4转流任务失败", zap.Error(create.Error))
		return errors.New(" 新增mix4转流任务失败")
	}
	return nil
}

// Update 更新
func (r *RoomMix4) Update() error {
	save := mysql.Instance.Save(&r)
	if save.Error != nil {
		log.L.Error("RoomMix4 更新失败", zap.Error(save.Error))
		return errors.New("RoomMix4 更新失败")
	}
	return save.Error

}

// Delete 删除
func (r *RoomMix4) Delete() error {
	save := mysql.Instance.Delete(&r)
	if save.Error != nil {
		log.L.Error("RoomMix4 删除失败", zap.Error(save.Error))
		return errors.New("RoomMix4 删除失败")
	}
	return save.Error

}

// GetById id查询
func GetById(id uint) (*RoomMix4, error) {
	var RoomMix4 RoomMix4
	mysql.Instance.First(&RoomMix4, id)

	if RoomMix4.ID == 0 {
		log.L.Sugar().Errorf("没有查询到该四合一画面任务 id:%d", id)
		return nil, errors.New("没有查询到该四合一画面任务")
	}

	return &RoomMix4, nil
}

func ListByIpAndRtspUrlsAndFfmpegSaveState(ip, rtspUrlMiddle, rtspUrlSmallOne, rtspUrlSmallTwo, rtspUrlSmallThree string, ffmpegSaveState int) (roomMix4s []RoomMix4) {
	mysql.Instance.Where("ip = ? and rtsp_url_middle = ? and rtsp_url_small_one = ? and rtsp_url_small_two = ? and rtsp_url_small_three = ? and ffmpeg_save_state = ?",
		ip, rtspUrlMiddle, rtspUrlSmallOne, rtspUrlSmallTwo, rtspUrlSmallThree, ffmpegSaveState).Find(&roomMix4s)
	return
}
