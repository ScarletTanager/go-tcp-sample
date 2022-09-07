package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/signal"
)

// Simple TCP echo server

func main() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("Fatal error: %s\n", err.Error())
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	log.Println("Listening on port 9999...")
	// defer listener.Close()

	go func() {
		<-sigchan
		log.Println("Shutting down listener...")
		listener.Close()
		log.Println("Exiting...")
		os.Exit(0)
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("ERROR: From Accept(): %s\n", err.Error())
		}

		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(conn)
	}
}
