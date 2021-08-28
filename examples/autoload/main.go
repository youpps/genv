package main

import (
	"fmt"
	"os"

	_ "github.com/youpps/genv/autoload"
)

func main() {
	fmt.Println(os.Getenv("KEY")) // 1234567890
}
