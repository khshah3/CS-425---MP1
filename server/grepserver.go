package main

import (
  "./grep"
  "net"
  "net/rpc"
  "net/http"
  "log"
)

// This server uses net/rpc magic to expose an object's methods over tcp
// The client can simply call the method as if it were local (almost)
func main() {
  myGrep := new(grep.Grep)
  // Expose the grep functionality over HTTP 
  rpc.Register(myGrep)
  rpc.HandleHTTP()

  // Start the server 
  port := ":1234"
  log.Println("start server on port:", port[1:])
  l, e := net.Listen("tcp", ":1234")
  if e != nil {
    log.Fatal("listen error:", e)
  }

  // TODO can we parallelize handling requests?
  for {
    http.Serve(l, nil)
  }
}
