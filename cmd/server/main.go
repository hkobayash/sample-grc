package main

import (
	"log"
	"net"

	. "github.com/hkobayash/sample-grc/internal/app"
	"github.com/hkobayash/sample-grc/internal/server"
)

func main() {
	app, err := NewApp("config.toml")
	if err != nil {
		log.Fatalf("Failed to initialize: %v", err)
	}

	l, err := net.Listen("tcp", app.Config.Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := server.NewServer(app)
	server.Serve(l)
}
