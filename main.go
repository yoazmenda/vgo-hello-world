package main

import "fmt"
import log "github.com/sirupsen/logrus"
func main(){
    fmt.Println("hello world")
    log.Info("this log is printed from logrus dependency")
}

