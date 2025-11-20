package main

import (
	"AstraFlow-go/internal/api"
	"AstraFlow-go/pkg/config"
	"fmt"
)

func main() {
	config.InitConfig()
	r := api.InitRouter()

	port := config.Cfg.Server.Port
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
