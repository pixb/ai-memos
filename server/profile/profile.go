package profile

// 配置启动服务器配置
type Profile struct {
	// 服务器启动模式，prod，dev，demo
	Mode string
	// 服务器监听地址
	Addr string
	// 服务器监听端口
	Port int
	// UNIXSock 是 IPC binding path，Override Addr and Port if set
	UNIXSock string
	// Data path
	Data string
	// DSN points to where memos stores its own data
	DSN string
	// Driver is the database driver
	// sqlite, mysql
	Driver string
	// Version is the current version of server
	Version string
	// InstanceURL is the url of your memos instance.
	InstanceURL string
}
