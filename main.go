package main

import (
	"fmt"
	"os"
)

func main() {
	envVar := os.Getenv("MY_VAR")
	fmt.Println(envVar)
	fmt.Println(os.Environ())
	secret, _ := os.ReadFile("/run/secret/api_token")
	fmt.Println(secret)
}
