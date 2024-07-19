package main

import (
	"log"

	"github.com/haormj/cyber/cmd/cyber/cmd"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	cmd.Execute()
}
