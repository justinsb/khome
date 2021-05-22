package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		fmt.Fprintf(os.Stdout, "Hello world\n")
		time.Sleep(time.Second)
	}
}
