package main

import(
  "net/rpc"
  "log"
  "./server/grep"
  "./serverlist"
  "os"
)

func remoteGrep(serverAddress, logPath, key, val string) {
  log.Println("dialing server:", serverAddress)
  client, err := rpc.DialHTTP("tcp", serverAddress)
  if err != nil {
    log.Fatal("dialing:", err)
  }

  // Give grep the path to the log file
  args := &grep.Args{logPath, key, val}
  var reply grep.Reply
  // Make the remote call (this is just a test)
  err = client.Call("Grep.Search", args, &reply)
  if err != nil {
    log.Fatal("grep search error:", err)
  }
  log.Printf("Search: %s:%s ==> %s", args.Key, args.Val, reply.Val)
}

func main() {
  // Make sure we have the right number of arguments
  if len(os.Args) != 3 {
    log.Panic("args:", "<keyQuery> <valQuery>")
  }
  // Get the command line arguments corresponding to key and value queries
  keyQuery, valQuery := os.Args[1], os.Args[2]

  // Get the configuration for all the servers running dgrep
  serverList := serverlist.FromFile("serverlist/testconfig.cfg")
  for i := range *serverList{
    //log.Println((*serverList)[server])
    server := (*serverList)[i]
    remoteGrep(server.Address, server.Logfile, keyQuery, valQuery)
  }
}
