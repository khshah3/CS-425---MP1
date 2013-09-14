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


  // Give grep the path to the log file
  defaultPath := "./sample_logs/machine.1.log"
  args := &grep.Args{defaultPath, "they", "ideas"}
  var reply grep.Reply
  // Make the remote call (this is just a test)
  err = client.Call("Grep.Search", args, &reply)
  if err != nil {
    log.Fatal("grep search error:", err)
  }
  log.Printf("Search: %s:%s ==> %s", args.Key, args.Val, reply.Val)
}
