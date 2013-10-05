package grep

import (
  "log"
  "fmt"
  "os/exec"
  "io/ioutil"
)

// The arguments to a remote call of grep
type Args struct {
  Filepath, Key, Val string
}

// The structure of grep's return value
type Reply struct {
  Val string
}

// Grep type, with the Search method exposed
type Grep int

// The method to actually grep local logs and return any matches
func (g *Grep) Search(args *Args, reply *Reply) error {
  log.Printf("search:file=%s, query=%s:%s", args.Filepath, args.Key, args.Val)

  // Build the grep regex to separate key/val searches (':' delimeter)
  query := fmt.Sprintf("^.*%s.*:.*%s.*$", args.Key, args.Val)
  fmt.Println(query)
  cmd := exec.Command("sh" , "-c" , "grep -r "+ query + " " + args.Filepath)

  // Get the command's stdout pipe so we can use the results of the call
  stdout, err := cmd.StdoutPipe()
  if err != nil {
    log.Fatal(err)
  }
  // Run the command
  if err := cmd.Start(); err != nil {
    log.Fatal(err)
  }

  // Use the command's stdout and convert to a reasonable return value
  // (we can't stream pipes over rpc)
  result, err := ioutil.ReadAll(stdout)
  reply.Val = string(result)

  log.Println("Finished:", "Sending results.")

  return nil
}


