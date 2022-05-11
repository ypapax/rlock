package main

import (
	"log"
	"sync"
)

func main(){
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	var rw = sync.RWMutex{}
	rw.RLock()
	log.Println("RLock")
	rw.Lock() // deadlock here
	log.Println("Lock")
}