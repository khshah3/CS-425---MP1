package serverlist

import (
  "os"
  "log"
  "bufio"
  "strings"
)

type ServerConfig struct {
  Address, Logfile string
}

type ServerList []*ServerConfig

func NewServerConfig(configLine string) *ServerConfig {
  fields := strings.Split(configLine, " ")
  sc := new(ServerConfig)
  sc.Address = fields[0]
  sc.Logfile = fields[1]
  return sc
}

// Read a file of address/logfile pairs on each line
// and create a ServerList from this information
func FromFile(filename string) *ServerList {
  sl := new(ServerList)
  file, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    *sl = append(*sl, NewServerConfig(scanner.Text()))
  }
  return sl
}


