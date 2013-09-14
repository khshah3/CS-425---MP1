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
  defaultPath := "/home/martin/school/cs425/mp1/sample_logs/machine.1.log"
  args := &grep.Args{defaultPath, "they", "idea"}
  var reply grep.Reply
  // Make the remote call (this is just a test)
  err = client.Call("Grep.Search", args, &reply)
  if err != nil {
    log.Fatal("grep search error:", err)
  }
  log.Printf("Search: %s:%s ==> %s", args.Key, args.Val, reply.Val)
}
