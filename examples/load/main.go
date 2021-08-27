package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ipsl/genv"
)

func main() {
	err := genv.Load(".env-custom")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getenv("KEY")) // 0987654321
}
