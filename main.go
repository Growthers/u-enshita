package main

import "fmt"

var (
	version  = "unknown"
	revision = "unknown"
)

func main() {
	fmt.Printf("mu-enshita\n(C) 2023 Growthers\n%v (%v)\n", version, revision)
}
