package main

import(
  "net/rpc"
  "log"
  "./server/grep"
)

func main() {
  serverAddress := "localhost"
  client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
  if err != nil {
    log.Fatal("dialing:", err)
  }

  // Make the remote call (this is just a test)
  args := &grep.Args{"hello", "world"}
  var reply grep.Reply
  err = client.Call("Grep.Search", args, &reply)
  if err != nil {
    log.Fatal("grep search error:", err)
  }
  log.Printf("Search: %s:%s ==> %s", args.Key, args.Val, reply.Val)
}
