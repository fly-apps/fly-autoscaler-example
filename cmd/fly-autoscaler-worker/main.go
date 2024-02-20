package main

import (
	"fmt"
	"time"
)

func main() {
	d := 30 * time.Second
	fmt.Printf("Running for %s...\n", d)
	time.Sleep(d)
	fmt.Println("...shutting down")
}
