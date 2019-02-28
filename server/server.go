package main

import (
	"io"
	"log"
	"net"
  "crypto/tls"
	"os"
)

// Handles TC connection and perform synchorinization:
// TCP -> Stdout and Stdin -> TCP
func tcp_con_handle(con net.Conn) {
	chan_to_stdout := stream_copy(con, os.Stdout)
	chan_to_remote := stream_copy(os.Stdin, con)
	select {
	case <-chan_to_stdout:
		log.Println("Remote connection is closed")
	case <-chan_to_remote:
		log.Println("Local program is terminated")
	}
}

// Performs copy operation between streams: os and tcp streams
func stream_copy(src io.Reader, dst io.Writer) <-chan int {
	buf := make([]byte, 4096)
	sync_channel := make(chan int)
	go func() {
		defer func() {
			if con, ok := dst.(*tls.Conn); ok {
				con.Close()
				log.Printf("Connection from %v is closed\n", con.RemoteAddr())
			}
			sync_channel <- 0 // Notify that processing is finished
		}()
		for {
			var nBytes int
			var err error
			nBytes, err = src.Read(buf)
			if err != nil {
				if err != io.EOF {
					log.Printf("Read error: %s\n", err)
				}
				break
			}
			_, err = dst.Write(buf[0:nBytes])
			if err != nil {
				log.Fatalf("Write error: %s\n", err)
			}
		}
	}()
	return sync_channel
}


func main() {
  log.Println("Starting server on :27015")
  log.SetFlags(log.Lshortfile)
  cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
  if err != nil{
    log.Println(err)
    return
  }
  config := &tls.Config{Certificates: []tls.Certificate{cer}}
  ln, err := tls.Listen("tcp", ":27015", config)
  if err != nil {
    log.Println(err)
    return
  }
  defer ln.Close()
  for {
    conn, err := ln.Accept()
    if err != nil{
      log.Println(err)
      continue
    }
    log.Println("Client connected ", conn.RemoteAddr())
    go tcp_con_handle(conn)
  }
}
