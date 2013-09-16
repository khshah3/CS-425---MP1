package main

import (
    "fmt"
    "bufio"
    "os"
    "bytes"
    "math/rand"
    "../logger"
    "log"
)

func main(){

    logs := new(logger.Logger)
    logs.GetLogFileName("../logger/logconfig.cfg")
    file, err := os.Open("random.txt")
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    logSize := logs.Size()
    count := 0
    size := 0
    for size < logSize {
        
        random_count := ( (rand.Int() % 9) + 2 )
        var key string
        var value bytes.Buffer

        
        for scanner.Scan() {
            
            if count == 0 {
                key = scanner.Text()
            }
            scanner.Scan()
            value.WriteString(scanner.Text()+ " ")
            count++
            if count == random_count {
                break
            }
        }
        count = 0
        size := logs.Log(key , value.String())
        fmt.Println(size)
        value.Reset()
    }
}
