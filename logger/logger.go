package logger

import (
    "log"
    "os"
)

type Logger struct {
  filename string
}

func NewLogger(filename string) *Logger {
  l := new(Logger)
  l.filename = filename
  return l
}

func (logs *Logger) Log(key string , value string)(int64){
    file, err := os.OpenFile(logs.filename, os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
    defer file.Close()
    if err != nil {
        log.Fatal(err)
    }
    file.WriteString(key + ":" + value + "\n")
    fileLength , err := file.Seek(0, os.SEEK_CUR)
    return fileLength
}



