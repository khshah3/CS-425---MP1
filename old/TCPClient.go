package main

import (
    "log"
    "net"
    "encoding/gob"
)

type P struct {
    M, N int64
}

func main() {

    log.Println("Starting client");
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal("Connection error", err)
    }
    encoder := gob.NewEncoder(conn)
    p := &P{1, 2}
    log.Println(p);
    encoder.Encode(p)
    conn.Close()
    log.Println("done");
}

