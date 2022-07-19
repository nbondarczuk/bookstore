package main

import (
	"os"

	"bookstore/server"
)

func main() {
	if err := server.Run(); err != nil {
		os.Exit(1)
	}
}
