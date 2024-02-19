:: 切换到当前路径
cd %~dp0

:: 编译
go build -ldflags "-s -w" -o ./bin/manager/mgr.exe ./ktManager/src/
go build -ldflags "-s -w" -o ./bin/monitor/mon.exe ./ktMonitor/src/

:: 压缩
upx.exe ./bin/manager/mgr.exe
upx.exe ./bin/monitor/mon.exe