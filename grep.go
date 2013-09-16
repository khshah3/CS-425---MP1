package main

import(
  "log"
  "./serverlist"
  "os"
  "fmt"
  "sync"
  "./dgrep"
)




func main() {
fmt.Println("Beginning dgrep seach")
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
    go dgrep.RemoteGrep(server.Address, server.Logfile, keyQuery, valQuery, &wg)
  }
  wg.Wait()
  log.Println("Done:", "All servers complete.")
}
