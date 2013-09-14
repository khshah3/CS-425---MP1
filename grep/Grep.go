package main

import (
  "log"
  "os"
  "os/exec"
  "io"
)


func main() {
  log.Println("Start grep test")
  grep("test")
  log.Println("Done.")
}

func grep(query string) string {
  // The file to grep
  var filename = "testfile.txt"

  // The grep command object
  cmd := exec.Command("grep", query, filename)

  // Get the command's stdout pipe so we can use the results of the call
  stdout, err := cmd.StdoutPipe()
  if err != nil {
    log.Fatal(err)
  }

  // Run the command
  if err := cmd.Start(); err != nil {
    log.Fatal(err)
  }

  // Redirect command's stdout to system stdout (so we can see it in console)
  // This blocks until it receives an error or EOF
  io.Copy(os.Stdout, stdout)

  // Dummy return value, eventually we'll do something with the actual output
  return "nothing"
}
