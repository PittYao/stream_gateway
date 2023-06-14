package httpclient

import (
	"errors"
	"github.com/PittYao/stream_gateway/components/gin/response"
	"github.com/PittYao/stream_gateway/components/log"
	"github.com/PittYao/stream_gateway/internal/dto"
	"github.com/guonaihong/gout"
	"go.uber.org/zap"
)

func StopHttpClient(url string) error {
	response := response.Response{}

	err := gout.
		// POST请求
		POST(url).
		// 打开debug模式
		Debug(true).
		// SetJSON设置http body为json
		// 同类函数有SetBody, SetYAML, SetXML, SetForm, SetWWWForm
		SetJSON(gout.H{}).
		// BindJSON解析返回的body内容
		// 同类函数有BindBody, BindYAML, BindXML
		BindJSON(&response).
		// 结束函数
		Do()

	// 判断错误
	if err != nil {
		log.L.Error("结束存储异常", zap.Error(err))
		return errors.New("结束存储异常")
	}

	if response.Code != 200 {
		log.L.Error("结束存储异常", zap.Any("response", response))
		return errors.New(response.Msg)
	}

	return err
}

func ListStreamHttpClient(url string) (error, response.Response) {
	response := response.Response{}

	err := gout.
		// POST请求
		POST(url).
		// 打开debug模式
		Debug(true).
		// SetJSON设置http body为json
		// 同类函数有SetBody, SetYAML, SetXML, SetForm, SetWWWForm
		SetJSON(gout.H{}).
		// BindJSON解析返回的body内容
		// 同类函数有BindBody, BindYAML, BindXML
		BindJSON(&response).
		// 结束函数
		Do()

	// 判断错误
	if err != nil {
		log.L.Error("查询异常", zap.Error(err))
		return errors.New("查询异常"), response
	}

	if response.Code != 200 {
		log.L.Error("查询异常", zap.Any("response", response))
		return errors.New(response.Msg), response
	}

	return err, response
}

func RebootHttpClient(url string) (error, response.Response) {
	response := response.Response{}
	err := gout.
		// POST请求
		POST(url).
		// 打开debug模式
		Debug(true).
		// SetJSON设置http body为json
		// 同类函数有SetBody, SetYAML, SetXML, SetForm, SetWWWForm
		SetJSON(gout.H{}).
		// BindJSON解析返回的body内容
		// 同类函数有BindBody, BindYAML, BindXML
		BindJSON(&response).
		// 结束函数
		Do()

	// 判断错误
	if err != nil {
		log.L.Error("重启异常", zap.Error(err))
		return errors.New("重启异常"), response
	}

	if response.Code != 200 {
		log.L.Error("重启异常", zap.Any("response", response))
		return errors.New(response.Msg), response
	}

	return err, response
}

func UpgradeHttpClient(url string, dto dto.UpgradeReq) error {
	response := response.Response{}
	err := gout.
		// POST请求
		POST(url).
		// 打开debug模式
		Debug(true).
		// SetJSON设置http body为json
		// 同类函数有SetBody, SetYAML, SetXML, SetForm, SetWWWForm
		SetJSON(gout.H{
			"fileUrl": dto.FileUrl,
		}).
		// BindJSON解析返回的body内容
		// 同类函数有BindBody, BindYAML, BindXML
		BindJSON(&response).
		// 结束函数
		Do()

	return err
}
