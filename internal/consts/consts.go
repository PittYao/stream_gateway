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
	Upgrade   = "/upgrade"
)

const (
	RunIngError          int = -1 // 运行异常
	RebootError          int = -2 // 重启任务失败
	RunIng               int = 1  // 正在运行
	CloseSuccess         int = 2  // 关闭成功
	RebootTaskButRunning int = 3  // 重启任务时发现已有任务在运行中
	BootClose            int = 4  // 开机关闭,已有相同任务重启
	RebootSuccess        int = 5  // 重启成功
)

const (
	TsFile       = "ts"
	FirstTsFile  = "video000.ts"
	TsFilePrefix = "video"
	EmptyTsFile  = ""
	M3u8File     = "playlist.m3u8"
)
