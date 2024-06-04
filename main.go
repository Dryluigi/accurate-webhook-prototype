package main

import (
	"accurate-webhook-prototype/internal/server"
)

func main() {
	s := server.NewServer()

	s.Run()
}
