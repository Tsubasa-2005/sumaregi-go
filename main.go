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
	fmt.Printf("scope: %s\n", accessTokenResult.Scope)
	fmt.Printf("token_type: %s\n", accessTokenResult.TokenType)
	fmt.Printf("expires_in: %d\n", accessTokenResult.ExpiresIn)
	fmt.Printf("access_token: %s\n", accessTokenResult.AccessToken)
}
