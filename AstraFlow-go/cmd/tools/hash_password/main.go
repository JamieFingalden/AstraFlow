package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// 用法：go run cmd/tools/hash_password/main.go 123456
func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run cmd/tools/hash_password/main.go <明文密码>")
		os.Exit(1)
	}

	plain := os.Args[1]
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("生成密码哈希失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(hashed))
}
