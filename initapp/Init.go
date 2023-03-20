package initapp

import (
	"github.com/flechazo-x/go_common/util/timex"
	"wechat_robot/global"
)

func Init() {
	timex.InitTime() //初始化时间种子
	global.InitSysConfig()
	InitRedis() //初始化Redis
}
