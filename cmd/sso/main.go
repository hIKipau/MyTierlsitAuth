package main

import (
	"fmt"
	"github.com/hIKipau/protoMyTierlistAuth/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger

	// TODO: init app

	// TODO: init gRPC-server app
}
