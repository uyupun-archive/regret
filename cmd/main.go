package main

import (
	"github.com/uyupun/regret/cmd/config"
)

func main() {
	router := newRouter()

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	router.Logger.Fatal(router.Start(":" + cfg.ListenAddress))
}
