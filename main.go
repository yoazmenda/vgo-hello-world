package main

import "fmt"
import log "github.com/sirupsen/logrus"

func main(){
    fmt.Println("Hell from main")
    Hello()
}

func Hello(){
    fmt.Println("hello from Hello functions")
    log.Info("this log is printed from logrus dependency")
}



