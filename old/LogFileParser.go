
package main

import (
	"fmt"
	"bufio"
	"log"
	"os/exec"
	"os"
  "strings"
  "bytes"
)

func main() {

        echo := exec.Command("echo","Terminal Input Test Successful")
        echo.Stdin = strings.NewReader("some input")
        var out bytes.Buffer
        echo.Stdout = &out
        err := echo.Run()
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Terminal Test: %s\n", out.String())
        logFile := findLogFile ("hellogo.go")
        result := grepFile (logFile)
        fmt.Printf("%s",result)

}


//Finds the given log file based on the input
func findLogFile(fileName string)(string){

	fmt.Println("Beginning File Search")
	find := exec.Command("sh", "-c", "find "+fileName)
    find.Stdin = strings.NewReader("some input")
    var out bytes.Buffer
    find.Stdout = &out
    err := find.Run()
    if err != nil {
        log.Fatal(err)
        fmt.Println("run")
    }
    fmt.Printf("Printing Log File Name: %s\n", out.String())
	return out.String()
}

//Find the grep file search result 
func grepFile(fileName string )(string) {
	
	find := exec.Command("sh" , "-c" ,"grep main "+fileName)
    find.Stdin = strings.NewReader("some input")
    var out bytes.Buffer
    find.Stdout = &out
    err := find.Run()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Printing Grep Results Name: %s\n", out.String())
	return out.String()
}

//GREPs the given file based on the input
func readFile(){

	file, err := os.Open("hellogo.go")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
    	fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
    	log.Fatal(err)
	}
}
func keyFind() {
	
}

func valueFind() {

}
