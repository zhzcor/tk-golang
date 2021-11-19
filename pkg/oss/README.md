# Aliyun OSS Go Callback Server

## 关于
aliyun_oss_callback_server.go 是 *Golang* 版本的上传回调服务器应用示例。

## 环境
- `Go 1.5` 及以上。

## 编译
### Windows
- 在命令行提示符下进入 `aliyun_oss_callback_server.go` 所在目录， 执行 `go build aliyun_oss_callback_server.go` 即可编译生成可执行文件 `aliyun_oss_callback_server.exe`。 

### Linux
- 在命令行下进入 `aliyun_oss_callback_server.go` 所在目录， 执行 `go build aliyun_oss_callback_server.go` 即可编译生成可执行文件 `aliyun_oss_callback_server`。

## 运行
### Windows
- 命令行提示符下执行 `aliyun_oss_callback_server.exe` 启动回调服务器，默为监听端口为 ***80***。 

### Linux
- 命令行下执行 `aliyun_oss_callback_server` 启动回调服务器， 默为监听端口为 ***80***。 

**提示**：
- 如果需要指定监听 *端口* 和 *IP*，请使用 `aliyun_oss_callback_server.exe <ip> <port>` 。

## 协议
- MIT。
