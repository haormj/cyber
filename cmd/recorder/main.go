package main

import (
	"log"

	"github.com/haormj/cyber/cmd/recorder/cmd"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	cmd.Execute()
}
