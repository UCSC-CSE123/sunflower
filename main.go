package main

import (
	"log"

	"github.com/UCSC-CSE123/sunflower/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
