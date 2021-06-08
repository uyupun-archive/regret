package main

import (
	"os"

	"github.com/uyupun/regret/cmd/config"
)

func main() {
	router := newRouter()

	config.Load()

	router.Logger.Fatal(router.Start(":" + os.Getenv("APP_PORT")))
}
