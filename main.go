package main

import (
	"log"

	"github.com/Sean-Miningah/newgit/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
