// Package common
/******************************************************************************
  作者: dmzn@163.com 2024-01-14 13:58:15
  描述: 处理 mqtt 协议
******************************************************************************/
package main

import (
	"github.com/dmznlin/znlib-go/znlib"
	mt "github.com/eclipse/paho.mqtt.golang"
)

// onMqttMessge 2024-01-14 13:59:47
/*
 参数: cli,mqtt链路
 参数: msg,mqtt消息内容
 描述: 接收broker下发的消息
*/
func onMqttMessge(cli mt.Client, msg mt.Message) {
	znlib.Info(string(msg.Payload()))
}
