package main

import (
	"log"

	"github.com/dzonerzy/goKIT"
)

func handleClient(info *goKIT.ClientInfo, kit *goKIT.KIT) {
	for {
		pkt := kit.Read()
		if pkt == nil {
			log.Fatalf("Read: %v", kit.Error())
		} else {
			if kit.IsDisconnect(pkt) {
				log.Printf("Client %d disconnected", info.ClientID())
				kit.NotifyDisconnect(info)
				break
			} else {
				log.Printf("Receive '%s' from client %d", string(pkt.Content()), info.ClientID())
			}
		}
	}
}

func main() {
	kit := new(goKIT.KIT)
	if kit.Init() {
		if kit.Bind(goKIT.KIT_DEFAULT_ID) {
			for {
				var info, client = kit.ListenAndAccept()
				if info == nil {
					log.Fatalf("ListenAndAccept: %v", kit.Error())
				} else {
					log.Printf("Client %d connected", info.ClientID())
					go handleClient(info, client)
				}
			}
		} else {
			log.Fatalf("Bind: %v", kit.Error())
		}
	}
}
