package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	exitchan := make(chan os.Signal, 2)
	signal.Notify(exitchan, os.Interrupt, os.Kill)

	go func() {
		for {
			select {
			case sig := <-exitchan:
				if sysSig, ok := sig.(syscall.Signal); ok {
					os.Exit(int(sysSig))
				}
			default:
			}
		}
	}()

	for {
	}
}
