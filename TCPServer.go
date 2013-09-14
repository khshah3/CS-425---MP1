package main

import (
    "net"
    "encoding/gob"
    "log"
)

type P struct {
    M, N int64
}
func handleConnection(conn net.Conn) {
    log.Println("Handling connection");
    dec := gob.NewDecoder(conn)
    p := &P{}
    dec.Decode(p)
    log.Printf("Data Received %+v",p);
}

func main() {
    log.Println("Starting server");
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        // handle error
    }
    for {
        conn, err := ln.Accept() // this blocks until connection or error
        if err != nil {
            // handle error
            continue
        }
        go handleConnection(conn) // a goroutine handles conn so that the loop can accept other connections
    }
}
