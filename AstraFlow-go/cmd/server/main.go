package main

import (
	"AstraFlow-go/internal/api"
	"AstraFlow-go/pkg/config"
	"fmt"
)

func main() {
	config.InitConfig()
	r := api.InitRouter()

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
