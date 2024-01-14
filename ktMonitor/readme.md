### 前置机程序
功能清单:
> 1. 启动时连接`ktMamanger`获取运行参数.
> 2. 使用modbus连接下位机PLC
> 3. 采集PLC数据,上传至`ktManager`.
> 4. 接收`ktManager`指令,写入PLC.

** mqttt **
> 1. 前置 -> 云端 称为`up`上行
> 2. 云端 -> 前置 称为`down`下载
> 3. 基地址: `/kongtiao/base`,用于 前置 - 云端交换基础参数
>    a./kongtiao/base/up: 用于向云端发起请求
>    b./kongtiao/base/down: 用于接收云端的数据
> 4. 前置的mqtt账户权限,对 `up` 只写不读,防止其它client查看上行数据;对`down`只读不写,防止向其它client发送恶意数据.
>