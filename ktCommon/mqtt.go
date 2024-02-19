// Package common
/******************************************************************************
  作者: dmzn@163.com 2024-01-14 15:34:54
  描述: 基于 mqtt 的业务协议
******************************************************************************/
package common

import (
	"time"
)

const (
	Param_Timeout = 3 * time.Second //命令运行超时
)

const (
	Cmd_front_getparam uint8 = iota //前置:获取运行参数
	Cmd_cloud_setparam              //云端: 设置运行参数

	Cmd_cloud_runscript //云端: 执行脚本
	Cmd_cound_runshell  //云端: shell命令
)
