package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

// Simple TCP echo client

func main() {
	conn, err := net.Dial("tcp", ":9999")
	if err != nil {
		log.Fatalf("Error establishing connection: %s\n", err.Error())
	}
	log.Println("Connection established...")
	defer conn.Close()

	// sigchan := make(chan os.Signal, 1)
	// signal.Notify(sigchan, os.Interrupt)

	bytesToWrite := []byte("foo")

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	// Write until we drop
	readBuf := make([]byte, 256)
	for {
		select {
		// case <-sigchan:
		// 	err := conn.Close()
		// 	if err != nil {
		// 		log.Fatalf("Error on Close(): %s\n", err.Error())
		// 	}
		// 	log.Println("Exiting...")
		// 	os.Exit(0)
		case <-ticker.C:
			log.Println("Writing...")
			written, err := conn.Write(bytesToWrite)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Got an error: %s\n", err.Error())
			}

			if written < len(bytesToWrite) {
				fmt.Fprintf(os.Stderr, "Too few bytes written (%d, should have been %d)\n", written, len(bytesToWrite))
			}
			log.Printf("Wrote %d bytes\n", written)

			bytesRead, err := conn.Read(readBuf)
			if err != nil {
				if err != io.EOF {
					log.Printf("ERROR on Read(): %s\n", err.Error())
				} else {
					log.Println("EOF received")
				}
			}
			log.Printf("Received %d byte\n", bytesRead)
		}
	}
}
