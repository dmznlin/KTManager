// Package main
/******************************************************************************
  作者: dmzn@163.com 2024-01-16 18:50:32
  描述: 空调云管家 - 云端服务
******************************************************************************/
package main

import (
	. "github.com/dmznlin/znlib-go/znlib"
)

// 初始化lib库
var _ = InitLib(nil, nil)

func main() {
	Mqtt.StartWithUtils(eventOnMQTT)
	//启动mqtt

	WaitSystemExit(func() error {
		Mqtt.Stop()
		return nil
	})
}
