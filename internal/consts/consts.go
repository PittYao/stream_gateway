package consts

const (
	LogFormatConsole = "console"
	LogFormatJson    = "json"

	EOFError    = "EOF"
	EOFErrorMsg = "没有传入body参数"
)

const (
	// 服务类型

	Single       string = "single"
	Mix3         string = "mix3"
	Mix4         string = "mix4"
	PublicSingle string = "publicSingle"
)

const (
	// 服务端口

	Mix3Port   = ":8007"
	Mix4Port   = ":8006"
	SinglePort = ":8005"
	PublicPort = ":8004"
	RtspPort   = "554"
)

const (
	Http        = "http://"
	Localhost   = "127.0.0.1"
	M3u8UrlPort = ":8880"
)

const (
	// 接口地址

	Start     = "/start"
	Stop      = "/stop"
	StopAll   = "/stopAll"
	RebootAll = "/rebootAll"
	GetAll    = "/getAll"
)

const (
	RunIngError        int = -1 // 运行异常
	RunNetworkError    int = -2 // 网络运行异常
	RunIng             int = 2  // 正在运行
	CloseSuccess       int = 4  // 关闭成功
	ReBoot             int = 6  // 定时重启
	RebootNetworkError int = 7  // 网络运行异常已重启
	Boot               int = 8  // 开机重启
)

const (
	TsFile       = "ts"
	FirstTsFile  = "video000.ts"
	TsFilePrefix = "video"
	EmptyTsFile  = ""
	M3u8File     = "playlist.m3u8"
)
