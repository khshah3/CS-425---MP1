package main

import(
  "net/rpc"
  "log"
  "./server/grep"
  "./serverlist"
  "os"
  "fmt"
  "sync"
)

func remoteGrep(serverAddress, logPath, key, val string, wg *sync.WaitGroup) {
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


func main() {
  // Make sure we have the right number of arguments
  if len(os.Args) != 3 {
    log.Panic("args:", "<keyQuery> <valQuery>")
  }
  // Get the command line arguments corresponding to key and value queries
  keyQuery, valQuery := os.Args[1], os.Args[2]

  // Get the configuration for all the servers running dgrep
  serverList := serverlist.FromFile("serverconfig.cfg")
  var wg sync.WaitGroup
  for i := range *serverList{
    server := (*serverList)[i]
    wg.Add(1)
    go remoteGrep(server.Address, server.Logfile, keyQuery, valQuery, &wg)
  }
  wg.Wait()
  log.Println("Done:", "All servers complete.")
}
