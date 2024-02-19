// Package common
/******************************************************************************
  作者: dmzn@163.com 2024-01-14 13:58:15
  描述: 处理 mqtt 协议
******************************************************************************/
package main

import (
	"github.com/dmznlin/znlib-go/znlib"
)

// eventOnMQTT 2024-01-16 14:40:56
/*
 参数: cmd,指令
 描述: 处理 mqtt 协议指令
*/
func eventOnMQTT(cmd *znlib.MqttCommand) error {
	cmd.SwitchUpDown(false)
	cmd.VerifyUse = true
	cmd.Data = "hello,word"
	cmd.SendCommand(cmd.Topic, 2)
	return nil
}
