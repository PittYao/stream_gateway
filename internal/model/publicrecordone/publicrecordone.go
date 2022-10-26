package publicrecordone

import (
	"errors"
	"github.com/PittYao/stream_gateway/components/log"
	"github.com/PittYao/stream_gateway/components/mysql"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type PublicRecordOne struct {
	gorm.Model
	RtspUrl                  string
	Ip                       string
	Port                     string
	SavePath                 string
	FfmpegTransformState     int
	FfmpegTransformCmd       string
	FfmpegTransformErrorMsg  string
	FfmpegTransformStartTime *time.Time // 转流开始时间
	FfmpegTransformCloseTime *time.Time // 转流结束时间
	FfmpegSaveState          int
	FfmpegSaveCmd            string
	FfmpegSaveErrorMsg       string
	FfmpegSaveStartTime      *time.Time // 存流开始时间
	FfmpegSaveCloseTime      *time.Time // 存流结束时间
	FfmpegStateLog           string     // 流运行日志
	RebootRootId             uint       //重启任务的根id
	RebootParentId           uint       //重启任务的父id
	M3u8Url                  string     //m3u8地址
	FileRecentTime           *time.Time // 最新生成文件的时间
	TsFile                   string     // 最新ts文件地址
	DisuseAt                 *time.Time // 淘汰的时间，过期的文件可以被删除

	RtmpUrl string `gorm:"-"` // 转换后的rtmp流地址

}

// --- orm --- //

// Add 插入单个流任务
func (r *PublicRecordOne) Add() error {
	create := mysql.Instance.Create(r)
	if create.Error != nil {
		log.L.Error("新增公区转流任务失败", zap.Error(create.Error))
		return errors.New(" 新增公区转流任务失败")
	}

	return nil
}

// Update 更新
func (r *PublicRecordOne) Update() error {
	save := mysql.Instance.Save(&r)
	if save.Error != nil {
		log.L.Error("更新失败", zap.Error(save.Error))
		return errors.New("更新失败")
	}
	return save.Error

}

// GetById id查询
func GetById(id uint) (*PublicRecordOne, error) {
	var publicRecordOne PublicRecordOne
	mysql.Instance.First(&publicRecordOne, id)

	if publicRecordOne.ID == 0 {
		log.L.Error("DB中没有查询到该公区单画面任务", zap.Uint("id", id))
		return nil, errors.New("DB中没有查询到该公区单画面任务")
	}

	return &publicRecordOne, nil
}

// Delete
func (r *PublicRecordOne) Delete() error {
	save := mysql.Instance.Delete(&r)
	if save.Error != nil {
		log.L.Error("公区 删除失败", zap.Error(save.Error))
		return errors.New("公区 删除失败")
	}
	return save.Error

}

func ListByIpAndRtspUrlsAndFfmpegSaveState(ip, rtspUrl string, ffmpegSaveState int) (publicRecordOnes []PublicRecordOne) {
	mysql.Instance.Where("ip = ? and rtsp_url = ? and ffmpeg_save_state = ?",
		ip, rtspUrl, ffmpegSaveState).Find(&publicRecordOnes)
	return
}
