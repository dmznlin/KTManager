// Package common
/******************************************************************************
  作者: dmzn@163.com 2024-01-14 13:58:15
  描述: 处理 mqtt 协议
******************************************************************************/
package main

import (
	"bytes"
	"common"
	"github.com/dmznlin/znlib-go/znlib"
	"os"
	"os/exec"
)

// mqttCommand 2024-01-16 14:40:56
/*
 参数: cmd,指令
 描述: 处理 mqtt 协议指令
*/
func eventOnMQTT(cmd *znlib.MqttCommand) error {
	switch cmd.Cmd {
	case common.Cmd_cloud_runscript, common.Cmd_cound_runshell:
		doRunCloudShellScript(cmd)
	}
	return nil
}

// getParamFromCloud 2024-01-22 18:03:06
/*
 描述: 获取云端参数
*/
func getParamFromCloud() bool {
	znlib.Info("连接云管家获取运行参数...")
	cmd := znlib.MqttUtils.NewCommand()
	cmd.Cmd = common.Cmd_front_getparam
	cmd.Timeout = common.Param_Timeout

	cmd = cmd.SendCommand("", znlib.MqttQosNone)
	if cmd == nil {
		znlib.Error("获取云端参数失败.")
		return false
	}

	znlib.Info(cmd.Data)

	return true
}

// doRunCloudShellScript 2024-01-22 18:08:52
/*
 参数: cmd,远程命令数据
 描述: 执行服务器发送的脚本
*/
func doRunCloudShellScript(cmd *znlib.MqttCommand) {
	script := znlib.FixPathVar("$path/tmp/")
	if !znlib.FileExists(script, true) {
		znlib.MakeDir(script)
	}

	caller := "mqtt.doRunCloudScript"
	buf, err := znlib.NewEncrypter(znlib.EncryptBase64_STD, nil).DecodeBase64([]byte(cmd.Data))
	if err != nil {
		znlib.ErrorCaller(err, caller)
		return
	}

	if znlib.Application.IsDebug {
		znlib.Info(string(buf))
	}

	if cmd.Cmd == common.Cmd_cloud_runscript { //执行脚本
		script = script + "script.sh"
		err = os.WriteFile(script, buf, 0755)
		if err != nil {
			znlib.ErrorCaller(err, caller)
			return
		}

		shell := exec.Command(script)
		err = shell.Start() //执行并立即返回
		if err != nil {
			znlib.ErrorCaller(err, caller)
		}
	}

	if cmd.Cmd == common.Cmd_cound_runshell { //shell命令
		var stdout, stderr bytes.Buffer
		shell := exec.Command("bash", "-c", string(buf))
		shell.Stdout = &stdout // 标准输出
		shell.Stderr = &stderr // 标准错误

		err = shell.Run() //执行并等待
		if err != nil {
			znlib.ErrorCaller(err, caller)
		}

		outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
		outStr = znlib.StrTrim(outStr)
		errStr = znlib.StrTrim(errStr)

		if outStr != "" {
			if errStr != "" {
				outStr = outStr + "\n" + errStr
			}
		} else {
			if errStr != "" {
				outStr = errStr
			} else {
				outStr = "执行完毕"
			}
		}

		result := znlib.MqttUtils.NewCommand()
		result.Cmd = common.Cmd_cloud_runscript
		result.Data = outStr

		result.Topic = cmd.Topic
		result.SwitchUpDown(true)
		result.SendCommand(result.Topic, 2)
	}

}
