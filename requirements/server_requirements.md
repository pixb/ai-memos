# Server Requirements
这是服务端的需求文档。用来创建一个笔记应用`Memos`的后端服务。

## 实现细节
- 后端服务使用`golang`语言实现.
- 入口文件路径为`bin/memos/main.go`
- 使用 `github.com/spf13/cobra` 库来构建命令行界面。

- 使用了 `github.com/spf13/viper` 库来管理配置。配置可以来自命令行标志、环境变量。
- main()函数程序入口。
- 定义`rootCmd`变量。
	- 程序的名称:`memos` 
	- 简短描述: `An open source, lightweight note-taking service. Easily capture and share your great thoughts.`
- 定义`greetingBanner`变量
	- 一个`Memos`的ASCII图形。
- init()函数定义
	- viper设置了配置项的默认值，
		  - `mode` 默认为 `"dev"`
		  - `driver` 默认为 `"sqlite"`
		  - `port` 默认为 `8081`
	- 定义了多个命令行标志 (flags),允许用户在启动时自定义这些配置。
		- mode，默认值: dev，description: mode of server, can be "prod" or "dev" or "demo"
		- addr, 默认值: "", description: address of server
		- port, default: 8081, description: port of server
		- unix-sock, default: "", description: path to the unix socket, overrides --addr and --port
		- data, default: "", description: data directory
		- driver, default: "sqlite", description: database driver
		- dsn, default: "", description: database source name(aka. DSN)
		- instance-url, default: "", description: the url of your memos instance
	- viper.BindPFlag,绑定所有的flag
	- 设置环境变量前缀MEMOS
	- 自动绑定所有环境变量
	- 绑定`instance-url`到`MEMOS_INSTANCE_URL`
- 定义printGeetings打印所有配置信息
- 当调用命令时，调用printGrettings信息

## 编译脚本

在创建一个编译脚本`scripts/build.sh`来编译服务端程序。
- 生成目录`build`
- 判断是否时windows系统，如果是，则使用`go build -o build/memos.exe`命令编译程序。
- 如果不是windows系统，则使用`go build -o build/memos`命令编译程序。
- 编译成功，提示用户编译成功。
- 提示用户执行`build/memos -mode  dev`命令查看帮助信息。

## 命令行打印示例

./build/memos --help
An open source, lightweight note-taking service. Easily capture and share your great thoughts.

Usage:
  memos [flags]

Flags:
      --addr string           address of server
      --data string           data directory
      --driver string         database driver (default "sqlite")
      --dsn string            database source name(aka. DSN)
  -h, --help                  help for memos
      --instance-url string   the url of your memos instance
      --mode string           mode of server, can be "prod" or "dev" or "demo" (default "dev")
      --port int              port of server (default 8081)
      --unix-sock string      path to the unix socket, overrides --addr and --port

./build/memos 
2025/05/10 09:20:55 WARN failed to find migration history in pre-migrate error="SQL logic error: no such table: migration_history (1)"
Development mode is enabled
DSN:  /Users/pix/dev/code/web/memos/build/memos_dev.db
---
Server profile
version: 0.24.3
data: /Users/pix/dev/code/web/memos/build
addr: 
port: 8081
unix-sock: 
mode: dev
driver: sqlite
---

███╗   ███╗███████╗███╗   ███╗ ██████╗ ███████╗
████╗ ████║██╔════╝████╗ ████║██╔═══██╗██╔════╝
██╔████╔██║█████╗  ██╔████╔██║██║   ██║███████╗
██║╚██╔╝██║██╔══╝  ██║╚██╔╝██║██║   ██║╚════██║
██║ ╚═╝ ██║███████╗██║ ╚═╝ ██║╚██████╔╝███████║
╚═╝     ╚═╝╚══════╝╚═╝     ╚═╝ ╚═════╝ ╚══════╝
Version 0.24.3 has been started on port 8081
---
See more in:
👉Website: https://usememos.com
👉GitHub: https://github.com/usememos/memos
---