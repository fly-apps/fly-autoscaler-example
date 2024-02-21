package main

import (
	"fmt"
	"time"
)

func main() {
	d := 10 * time.Second

	fmt.Printf("Running for %s...\n", d)
	time.Sleep(d)
	fmt.Println("...shutting down")
}
