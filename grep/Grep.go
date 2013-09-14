package main

import (
  "log"
  "os"
  "os/exec"
  "io"
  "bufio"
  "fmt"
  "strings"
  "regexp"
)


func main() {
  log.Println("Start grep test")
  grep("test")
  regexSearch("2013", "fizz")
  log.Println("Done.")
}

func grep(query string) string {
  // The file to grep
  var filename = "testlog.txt"
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


func regexSearch(keyQuery, valQuery string) string {
  filename := "testlog.txt"
  file, err := os.Open(filename)
  if err != nil { log.Fatal(err) }
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    fields := strings.SplitN(scanner.Text(), " ", 3)
    key := strings.Join(fields[:2], " ")
    val := fields[2]
    fmt.Println(key)
    fmt.Println(val)
  }

  return ""

}
