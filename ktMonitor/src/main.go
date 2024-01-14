// Package main
/******************************************************************************
  作者: dmzn@163.com 2024-01-14 13:52:17
  描述: 空调云管家 - 前置机程序
******************************************************************************/
package main

import (
	"github.com/dmznlin/znlib-go/znlib"
	"ktmanager/ktCommon"
)

// 初始化znlib-go基础库
var _ = znlib.InitLib(nil, nil)

func main() {
	err := znlib.Mqtt.Start(onMqttMessge)
	if err != nil {
		znlib.Error("启动mqtt服务失败.", znlib.LogFields{"err": err})
		return
	}

	znlib.Info("连接云管家获取运行参数...")
	mc := common.NewMqttCommand()
	znlib.Info(mc)

	znlib.WaitSystemExit(func() error {
		znlib.Mqtt.Stop()
		return nil
	})
}
