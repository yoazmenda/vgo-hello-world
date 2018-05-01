package vgohelloword

import "fmt"
import log "github.com/sirupsen/logrus"

func Hello(){
    fmt.Println("hello from Hello functions")
    log.Info("this log is printed from logrus dependency")
}



