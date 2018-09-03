package service

import (
	"flag"
	"bank.explorer/config"
	"bank.explorer/common"
	"bank.explorer/util/logger"
)

func ConfigInit() {
	flag.StringVar(&config.CurUser, "u", "", "")
	flag.StringVar(&config.JobList, "l", "", "jobList")
	flag.Parse()

	// 获取本机IP
	localIp, err := common.GetLocalIp()
	if err != nil {
		panic("GetLocalIp Err:" + err.Error())
	}
	config.LocalIp = localIp

	// log init
	logger.SetConsole(false)
	logger.SetLevel(logger.INFO)
	logger.SetRollingDaily(common.GetLogPath())
}
