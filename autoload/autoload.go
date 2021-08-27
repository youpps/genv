package autoload

import (
	"log"

	"github.com/ipsl/genv"
)

func init() {
	err := genv.Load()
	if err != nil {
		log.Fatal("Error in genv initialization: " + err.Error())
	}
}
