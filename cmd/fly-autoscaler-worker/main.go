package main

import (
	"fmt"
	"time"
)

func main() {
	d := 5 * time.Minute

	fmt.Printf("Running for %s...\n", d)
	time.Sleep(d)
	fmt.Println("...shutting down")
}
