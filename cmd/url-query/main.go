package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	httpServer, cleanup, err := wireUp()
	if err != nil {
		panic(err)
	}

	handleInterruptSignal(func(c chan os.Signal) {
		<-c
		fmt.Println("shutting down...")
		cleanup()
	})

	if err := httpServer.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("http server terminated: %s", err.Error()))
	}

	fmt.Printf("shut down completed!")
}

func handleInterruptSignal(callback func(c chan os.Signal)) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go callback(c)
}
