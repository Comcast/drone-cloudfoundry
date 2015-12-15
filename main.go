package main

import (
	"fmt"
)

var (
	build     string
	buildDate string
)

func main() {
	fmt.Printf("Drone Cloud Foundry Plugin built at %s\n", buildDate)
}
