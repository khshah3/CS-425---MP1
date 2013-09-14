package grep

import (
  "log"
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
  reply.Val = args.Key
  return nil
}


