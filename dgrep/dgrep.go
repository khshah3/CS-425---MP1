package dgrep

import(
  "net/rpc"
  "log"
  "../server/grep"
  "fmt"
  "sync"
)

func RemoteGrep(serverAddress, logPath, key, val string, wg *sync.WaitGroup) {
  // Tell the WaitGroup we're done after the function returns
  defer wg.Done()

  log.Println("dialing server:", serverAddress)
  client, err := rpc.DialHTTP("tcp", serverAddress)
  if err != nil {
    log.Println("dialing:", err)
    return
  }

  // Give grep the path to the log file
  args := &grep.Args{logPath, key, val}
  var reply grep.Reply
  // Make the remote call (this is just a test)
  err = client.Call("Grep.Search", args, &reply)
  if err != nil {
    log.Fatal("grep search error:", err)
  }
  log.Printf("Search: %s:%s", args.Key, args.Val)
  fmt.Print(reply.Val)
}


