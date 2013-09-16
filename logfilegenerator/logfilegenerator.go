package main

import (
    "bufio"
    "os"
    "../logger"
    "log"
    "math/rand"
    "bytes"
    "strconv"
)

    func getRandomText() ([]string ,int){
    
        file, err := os.Open("random.txt")
        if err != nil {
            log.Fatal(err)
        }
        count := 0
        var words []string

        scanner := bufio.NewScanner(file)
        scanner.Split(bufio.ScanWords)
        for scanner.Scan(){
            words = append (words , scanner.Text() )
            count++
        }
        
        return words,count

    }
func main(){
    
    if len(os.Args) != 3 {
        log.Panic("args:", "<logLocation> <size>")
    }

    logs := new(logger.Logger)
    logs.SetLogFileName(os.Args[1])    
    
   
    words,count := getRandomText()
    tempLogSize,err := strconv.Atoi (os.Args[2])
    if err != nil {
            log.Fatal(err)
    }
    logSize := int64 (tempLogSize)

        
    var key string
    var value bytes.Buffer
    var size int64
    size = 0
    
    //Store rare frequent and somewhat frequent values , 0:rare 1:frequent 2:somewhat frequent
   
    var moderate,frequent float64
    moderate = float64(logSize)*0.05
    frequent = float64(logSize)*0.2
    
    specialKey := []string{"RARE" , "MODERATE" , "FREQUENT"}
    specialCount := []int64{1 ,int64(moderate), int64(frequent)}
    occurrence := []int64 { (logSize / specialCount[0]) -1 , (logSize / specialCount[1]) -1 , (logSize / specialCount[2]) -1}
    

    loop_count := 0
    spec_count := []int64{0,0,0}
    
    for size < logSize {
        
        random_count := ( (rand.Int() % 9) + 1)
        
        key = words[loop_count % count]
        loop_count++

        if spec_count[1] == occurrence[1] {
            key = specialKey[1]
            spec_count[1] = 0
        }
        
        if spec_count[0] == occurrence[0] {
            key = specialKey[0]
            spec_count[0] = 0
        }
        
        for j := 0; j<random_count; j++ {
            value.WriteString(words[loop_count % count]+ " ")
                
                if spec_count[2] == occurrence[2] {
                value.WriteString(specialKey[2]+ " ")
                spec_count[2] = 0
                }
            loop_count++
            
        }
        spec_count[0]++
        spec_count[1]++
        spec_count[2]++
        size = logs.Log(key , value.String())
        value.Reset()
    }
}
