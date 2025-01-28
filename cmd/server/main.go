package main

import (
	"flag"
	"log"

	"github.com/hewenyu/translate/common"
	"github.com/hewenyu/translate/server"
)

var (
	configFile = flag.String("config", "config.yaml", "配置文件路径")
	addr       = flag.String("addr", ":1188", "服务器监听地址")
)

func main() {
	flag.Parse()

	// 加载配置
	cfg, err := common.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 创建服务器
	srv, err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("创建服务器失败: %v", err)
	}

	// 运行服务器
	log.Printf("服务器启动在 %s", *addr)
	if err := srv.Run(*addr); err != nil {
		log.Fatalf("服务器运行失败: %v", err)
	}
}
