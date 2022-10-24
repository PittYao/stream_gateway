package components

import (
	"github.com/PittYao/stream_gateway/components/config"
	"github.com/PittYao/stream_gateway/components/log"
	"github.com/PittYao/stream_gateway/components/mysql"
	"github.com/PittYao/stream_gateway/components/nginx"
	"github.com/PittYao/stream_gateway/components/swagger"
)

func Init() {
	config.Load()
	log.Load()
	swagger.Load()
	mysql.Load()
	nginx.Load()
	log.L.Info("项目初始化配置完成")
}
