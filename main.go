package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	ch := make(chan os.Signal, 100)
	signal.Notify(ch)

	fmt.Println("Signal monitor started. Waiting for signals...")

	for sig := range ch {
		fmt.Printf("Received signal: %s (%d)\n", sig.String(), sig)
	}
}
