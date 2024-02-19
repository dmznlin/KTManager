// Package main
/******************************************************************************
  作者: dmzn@163.com 2024-01-14 13:52:17
  描述: 空调云管家 - 前置机程序
******************************************************************************/
package main

import (
	"flag"
	"fmt"
	. "github.com/dmznlin/znlib-go/znlib"
)

// 初始化lib基础库
var _ = InitLib(nil, nil)

func main() {
	var pass string
	flag.StringVar(&pass, "pass", "", "生成DB、MQTT的DES密码")
	flag.Parse()

	if pass != "" {
		buf, err := NewEncrypter(EncryptDES_ECB,
			[]byte(DefaultEncryptKey)).Encrypt([]byte(pass), true)
		if err == nil {
			fmt.Println(string(buf))
		} else {
			Error(err)
		}

		return
	}

	//  ---------------------------------------------------------------------------
	Mqtt.StartWithUtils(eventOnMQTT)
	//启动mqtt

	WaitSystemExit(func() error {
		Mqtt.Stop()
		return nil
	})
}
