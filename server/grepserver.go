package main

import (
  "./grep"
  "net"
  "net/rpc"
  "net/http"
  "log"
)

func main() {
  // Expose the grep functionality over HTTP 
  rpc.Register(new(grep.Grep))
  rpc.HandleHTTP()
  l, e := net.Listen("tcp", ":1234")
  if e != nil {
    log.Fatal("listen error:", e)
  }

  // TODO is this the best way to do this?
  for {
    http.Serve(l, nil)
  }
}
