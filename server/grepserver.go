package main

import (
  "./grep"
  "net"
  "net/rpc"
  "net/http"
  "log"
  "os"
)

func startServer(port, logfile string) {
  // Expose the grep functionality over HTTP 
  myGrep := new(grep.Grep)
  rpc.Register(myGrep)
  rpc.HandleHTTP()

  // Start the server 
  log.Println("start server on port:", port)
  l, e := net.Listen("tcp", ":" + port)
  if e != nil {
    log.Fatal("listen error:", e)
  }

  // TODO can we parallelize handling requests?
  for {
    http.Serve(l, nil)
  }
}

// This server uses net/rpc magic to expose an object's methods over tcp
// The client can simply call the method as if it were local (almost)
func main() {
  if len(os.Args) != 3 {
    log.Panic("args:", "<port> <logfile>")
  }
  port, logfile := os.Args[1], os.Args[2]
  startServer(port, logfile)
}
