package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rhoeasy/concurency-in-go/pool"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	exitchan := make(chan os.Signal, 2)
	signal.Notify(exitchan, os.Interrupt, os.Kill)

	go func() {
		for {
			select {
			case sig := <-exitchan:
				log.Printf("app is canceled with signal[%v]", sig)
				cancel()
				time.Sleep(1 * time.Second)
				if sysSig, ok := sig.(syscall.Signal); ok {
					os.Exit(int(sysSig))
				}

			default:
			}
		}
	}()
	p := pool.NewPool(3, 2)
	go p.Run(ctx)
	for c := 'A'; c <= 'Z'; c++ {
		p.SubmitTask(string(c))
	}
	select {}
}
