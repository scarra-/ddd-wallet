package main

import (
	"log"

	"github.com/aadejanovs/wallet/internal/infrastructure"
)

func main() {
	server := infrastructure.Setup()
	log.Fatal(server.Listen("0.0.0.0:80"))
}
