package logger

import (
	
    "bufio"
    "log"
    "os"
    "strings"
    "strconv"
)

type Logger struct {
 	key ,value , filename string
    size int64
}

func (logs *Logger) GetLogFileName(filename string){
  file, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }
  scanner := bufio.NewScanner(file)
  scanner.Scan()
  fields := strings.Split(scanner.Text(), " ")
  logs.filename = fields[0]
  size ,err := strconv.Atoi(fields[1])
  logs.size = int64(size)
  file.Close()

}

func (logs *Logger) Size()(size int64){
        return logs.size
    }

func (logs *Logger) Log(key string , value string)(int64){

    file, err := os.OpenFile(logs.filename , os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
    if err != nil {
        log.Fatal(err)
    }
    file.WriteString(key + ":" + value + "\n")
    fileLength , err := file.Seek(0, os.SEEK_CUR)  
    file.Close() 
    return fileLength
 
}



