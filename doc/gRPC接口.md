grpcurl 测试命令

```shell

# 展示所有service
grpcurl -plaintext  localhost:9115 list
# 展示所有接口信息
grpcurl -plaintext  localhost:9115 describe device_watcher.Scanner 
# 请求所有网卡设备
grpcurl -plaintext  localhost:9115  device_watcher.Scanner/ScanNicDevice

```