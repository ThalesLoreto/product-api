package main

import (
	"fmt"

	"github.com/ThalesLoreto/product-api/configs"
)

func main() {
	cfg, _ := configs.LoadConfig(".")
	fmt.Println(cfg.DBDriver)
}
