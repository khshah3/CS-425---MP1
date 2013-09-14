package main

// Generate a log file with certain terms appearing more frequently than others
// Run it and redirect output to a file using ./loggen &> <filename>
// Terminate the program once the desired length has been reached (hack)

import (
  "log"
)

func main() {
  for i := 0; ; i++ {
    fizz, buzz := "", ""
    if i % 3 == 0 { fizz = "fizz" }
    if i % 5 == 0 { buzz = "buzz" }
    log.Printf("This is the %dth log message. %s%s", i, fizz, buzz)
  }
}

