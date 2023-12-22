package mall

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"mall/api/router"
	"mall/internal/core"
	"mall/internal/logger"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

var webCmd = &cobra.Command{
	Use: "web",
	Run: startWebServer,
}

var config string

func init() {
	rootCmd.AddCommand(webCmd)
	webCmd.Flags().StringVarP(&config, "config", "c", "", "config file path")
}

func startWebServer(cmd *cobra.Command, args []string) {
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	err := core.InitConfig(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = logger.InitLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 初始化gin
	engine := gin.New()
	// 路由
	router.RegisterRouter(engine)
	// 初始化server
	server := initServer(engine)
	// 启动Web服务
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(fmt.Sprintf("ListenAndServe error: %v", err))
	}
}

func initServer(handler http.Handler) *http.Server {
	server := &http.Server{
		Addr:         core.GlobalConfig.Server.Addr,
		Handler:      handler,
		ReadTimeout:  core.GlobalConfig.Server.ReadTimeout,
		WriteTimeout: core.GlobalConfig.Server.WriteTimeout,
		IdleTimeout:  core.GlobalConfig.Server.IdleTimeout,
	}

	return server
}
