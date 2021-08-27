package main

import (
	"fmt"
	"os"

	_ "github.com/ipsl/genv/autoload"
)

func main() {
	fmt.Println(os.Getenv("KEY")) // 1234567890
}
