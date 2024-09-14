package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	LoadEnv()

	accessTokenResult, err := GetAccessToken()
	if err != nil {
		log.Fatalf("Failed to get access token: %v", err)
	}

	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Printf("access_token: %s\n", accessTokenResult)
}
