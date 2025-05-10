package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const greetingBanner = `
███╗   ███╗███████╗███╗   ███╗ ██████╗ ███████╗
████╗ ████║██╔════╝████╗ ████║██╔═══██╗██╔════╝
██╔████╔██║█████╗  ██╔████╔██║██║   ██║███████╗
██║╚██╔╝██║██╔══╝  ██║╚██╔╝██║██║   ██║╚════██║
██║ ╚═╝ ██║███████╗██║ ╚═╝ ██║╚██████╔╝███████║
╚═╝     ╚═╝╚══════╝╚═╝     ╚═╝ ╚═════╝ ╚══════╝
`

var rootCmd = &cobra.Command{
	Use:   "memos",
	Short: "An open source, lightweight note-taking service. Easily capture and share your great thoughts.",
	Run: func(cmd *cobra.Command, args []string) {
		printGreetings()
		// 在这里添加服务启动逻辑
		fmt.Println("Version 0.24.3 has been started on port", viper.GetInt("port"))
		fmt.Println("---")
		fmt.Println("See more in:")
		fmt.Println("👉Website: https://usememos.com")
		fmt.Println("👉GitHub: https://github.com/usememos/memos")
		fmt.Println("---")
	},
}

func init() {
	viper.SetDefault("mode", "dev")
	viper.SetDefault("driver", "sqlite")
	viper.SetDefault("port", 8081)

	rootCmd.PersistentFlags().String("mode", "dev", `mode of server, can be "prod" or "dev" or "demo"`)
	rootCmd.PersistentFlags().String("addr", "", "address of server")
	rootCmd.PersistentFlags().Int("port", 8081, "port of server")
	rootCmd.PersistentFlags().String("unix-sock", "", "path to the unix socket, overrides --addr and --port")
	rootCmd.PersistentFlags().String("data", "", "data directory")
	rootCmd.PersistentFlags().String("driver", "sqlite", "database driver")
	rootCmd.PersistentFlags().String("dsn", "", "database source name(aka. DSN)")
	rootCmd.PersistentFlags().String("instance-url", "", "the url of your memos instance")

	viper.BindPFlag("mode", rootCmd.PersistentFlags().Lookup("mode"))
	viper.BindPFlag("addr", rootCmd.PersistentFlags().Lookup("addr"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("unix-sock", rootCmd.PersistentFlags().Lookup("unix-sock"))
	viper.BindPFlag("data", rootCmd.PersistentFlags().Lookup("data"))
	viper.BindPFlag("driver", rootCmd.PersistentFlags().Lookup("driver"))
	viper.BindPFlag("dsn", rootCmd.PersistentFlags().Lookup("dsn"))
	viper.BindPFlag("instance-url", rootCmd.PersistentFlags().Lookup("instance-url"))

	viper.SetEnvPrefix("MEMOS")
	viper.AutomaticEnv()
	viper.BindEnv("instance-url", "MEMOS_INSTANCE_URL")
}

func printGreetings() {
	if viper.GetString("mode") == "dev" {
		fmt.Println("Development mode is enabled")
	}
	// 实际应用中，DSN的生成会更复杂，这里仅为示例
	dataDir := viper.GetString("data")
	if dataDir == "" {
		dataDir, _ = os.Getwd()      // 默认为当前目录
		dataDir = dataDir + "/build" // 需求中DSN指向 build 目录下的db文件
	}
	dsn := viper.GetString("dsn")
	if dsn == "" {
		dsn = dataDir + "/memos_" + viper.GetString("mode") + ".db"
	}
	fmt.Println("DSN:", dsn)
	fmt.Println("---")
	fmt.Println("Server profile")
	fmt.Println("version: 0.24.3") // 版本号暂时硬编码
	fmt.Println("data:", dataDir)
	fmt.Println("addr:", viper.GetString("addr"))
	fmt.Println("port:", viper.GetInt("port"))
	fmt.Println("unix-sock:", viper.GetString("unix-sock"))
	fmt.Println("mode:", viper.GetString("mode"))
	fmt.Println("driver:", viper.GetString("driver"))
	fmt.Println("---")
	fmt.Println(greetingBanner)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
