package main

import (
	"fmt"
	"github.com/flechazo-x/go_common/logx"
	"wechat_robot/initapp"
	"wechat_robot/module/message"
	"wechat_robot/static"
)

func main() {
	logx.InitMultiple("config/logs_config.json") // 初始化多个日志句柄
	defer logx.Destruct()                        // 析构资源
	initapp.Init()                               //初始化app

	fmt.Println(message.GetMessage(static.Everyday))
	//fmt.Println(message.GetMessage(static.HotSearch))
}
