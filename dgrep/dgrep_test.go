package dgrep

import(
	"testing"
    "../serverlist"
    "sync"
)

func BenchmarkDgrepModerate(b *testing.B){

for j := 0; j < b.N; j++ {
    keyQuery, valQuery := "MODERATE","*"

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

func BenchmarkDgrepFrequent(b *testing.B){

for j := 0; j < b.N; j++ {
    keyQuery, valQuery := "*","FREQUENT"

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

func BenchmarkDgrepRare(b *testing.B){

for j := 0; j < b.N; j++ {
    keyQuery, valQuery := "RARE","*"

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
