package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		panic("Usage: program <buffer_size>")
	}

	bufferSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Invalid buffer size: %s must be an integer", os.Args[1]))
	}
	if bufferSize <= 0 {
		panic("Buffer size must be greater than 0")
	}

	ch := make(chan os.Signal, bufferSize)
	signal.Notify(ch, syscall.SIGUSR1)

	pid := os.Getpid()
	ppid := os.Getppid()
	pgid, _ := syscall.Getpgid(pid)
	fmt.Printf("PID: %d, PPID: %d, PGID: %d", pid, ppid, pgid)
	fmt.Printf("Signal monitor started with buffer size %d. Waiting for SIGUSR1...\n", bufferSize)
	fmt.Println("Press Ctrl+C or send SIGTERM to exit gracefully")

	for sig := range ch {
		fmt.Printf("Received signal: %s (%d)\n", sig.String(), sig)
	}
}
