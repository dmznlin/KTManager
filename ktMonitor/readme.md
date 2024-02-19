### 前置机程序
功能清单:
> 1. 启动时连接`ktMamanger`获取运行参数.
> 2. 使用modbus连接下位机PLC
> 3. 采集PLC数据,上传至`ktManager`.
> 4. 接收`ktManager`指令,写入PLC.

### mqttt 
> 1. 前置 -> 云端 称为`up`上行
> 2. 云端 -> 前置 称为`down`下载
> 3. 通道: `/ktmanager/tunnel`,用于 前置 - 云端交换数据\
>    a./ktmanager/tunnel/up/0101: 设备0101向云端发起请求\
>    b./ktmanager/tunnel/down/0101: 云端向设备0101发送数据
> 4. 前置的mqtt账户权限,对 `up` 只写不读,防止其它client查看上行数据;对`down`只读不写,防止向其它client发送恶意数据.

### 配置参数
1. mqtt.clientID: 用于识别mqtt连接,每个实例不能重复.如果两个id重复,会引起mqtt断开异常.
```  
msg="znlib.mqtt.connect: mqtts://192.168.2.202:1883"
msg="znlib.mqtt.subscribe: map[/ktmanager/+/up/+:0]"
msg="znlib.mqtt.reconnect_broker."
```
2. 