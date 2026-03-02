package main

import (
	"log"

	"github.com/DanielChachagua/anvil/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
