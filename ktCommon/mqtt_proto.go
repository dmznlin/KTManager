// Package common
/******************************************************************************
  作者: dmzn@163.com 2024-01-14 15:34:54
  描述: 基于 mqtt 的业务协议
******************************************************************************/
package common

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/dmznlin/znlib-go/znlib"
)

const (
	cmd_encryptkey = "mqtt.key" //加密密钥

	Cmd_front_getparam uint8 = iota //前置:获取运行参数
	Cmd_cloud_setparam              //云端: 设置运行参数
)

// MqttCommand 命令结构
type MqttCommand struct {
	Serial string `json:"serial"`
	Sender string `json:"sender"`
	Cmd    uint8  `json:"cmd"`
	Data   string `json:"data"`
	Verify string `json:"verify"`
}

// NewMqttCommand 2024-01-14 16:33:15
/*
 描述: 初始化一个命令
*/
func NewMqttCommand() *MqttCommand {
	no, _ := znlib.SerialID.NextStr(false)
	//业务序列号

	return &MqttCommand{
		Serial: no,
		Sender: znlib.Mqtt.Options.ClientID,
		Cmd:    0,
		Data:   "",
		Verify: "",
	}
}

// GetVerify 2024-01-14 15:48:28
/*
 描述: 计算验证信息
*/
func (mc *MqttCommand) GetVerify() string {
	mc.Verify = cmd_encryptkey
	data, err := json.Marshal(mc)
	//占位后生成json字符串

	if err != nil {
		znlib.Error("mqtt_proto: GetVerify", znlib.LogFields{"err": err})
		return ""
	}

	mc.Verify = fmt.Sprintf("%x", md5.Sum(data))
	return mc.Verify
}

// VerifyData 2024-01-14 16:19:13
/*
 参数: caller,调用方
 描述: 验证caller的协议是否有效
*/
func (mc MqttCommand) VerifyData(caller string) bool {
	origin := mc.Verify
	ok := origin == mc.GetVerify()
	if !ok {
		znlib.Warn(fmt.Sprintf("mqtt_proto.VerifyData: %s > invalid.", caller))
	}

	return ok
}
