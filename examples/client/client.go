package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dzonerzy/goKIT"
)

func main() {
	kit := new(goKIT.KIT)
	if kit.Init() {
		if kit.Connect(goKIT.KIT_DEFAULT_ID) {
			log.Println("Connected")
			counter := 1
			for {
				if counter >= 10 {
					break
				}
				if kit.Write(fmt.Sprintf("Hello world %d", counter)) {
					log.Println("Data sent")
					counter++
					time.Sleep(1 * time.Second)
				} else {
					log.Fatalf("Write: %v", kit.Error())
				}
			}
			if kit.Disconnect() {
				log.Println("Disconnected")
				return
			}
		} else {
			log.Fatalf("Connect: %v", kit.Error())
		}
	}
}
