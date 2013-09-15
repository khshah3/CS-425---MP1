package main

import(
  "net/rpc"
  "log"
  "./server/grep"
  "os"
)

func main() {
  // Make sure we have the right number of arguments
  if len(os.Args) != 3 {
    log.Panic("args:", "<keyQuery> <valQuery>")
  }
  // Get the command line arguments corresponding to key and value queries
  keyQuery, valQuery := os.Args[1], os.Args[2]


  serverAddress := "localhost"
  port := ":1234"
  log.Println("dialing server:", serverAddress + port)
  client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
  if err != nil {
    log.Fatal("dialing:", err)
  }


  // Give grep the path to the log file
  defaultPath := "./sample_logs/machine.1.log"
  args := &grep.Args{defaultPath, keyQuery, valQuery}
  var reply grep.Reply
  // Make the remote call (this is just a test)
  err = client.Call("Grep.Search", args, &reply)
  if err != nil {
    log.Fatal("grep search error:", err)
  }
  log.Printf("Search: %s:%s ==> %s", args.Key, args.Val, reply.Val)
}
