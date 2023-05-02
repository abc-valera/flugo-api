package main

import (
	"log"

	"github.com/abc-valera/flugo-api/internal/server"
)

func main() {
	log.Fatal(server.RunServer())
}
