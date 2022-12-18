package main

import (
	"fmt"

	"github.com/giovane-aG/goexpert/9-APIs/configs"
)

func main() {
	config := configs.LoadConfig(".")
	fmt.Println(config.DBDriver)
}
