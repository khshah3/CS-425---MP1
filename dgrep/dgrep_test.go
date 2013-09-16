package dgrep

import(
	"testing"
    "../serverlist"
    "sync"
)

func BenchmarkDgrep(b *testing.B){

for j := 0; j < b.N; j++ {
    keyQuery, valQuery := "this","it"

    // Get the configuration for all the servers running dgrep
    serverList := serverlist.FromFile("../serverconfig.cfg")
    var wg sync.WaitGroup
    for i := range *serverList{
        server := (*serverList)[i]
        wg.Add(1)
        go RemoteGrep(server.Address, server.Logfile, keyQuery, valQuery, &wg)
    }
    wg.Wait()
    }

}
