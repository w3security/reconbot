package main

import (
	"github.com/w3security/reconbot/pkg/hydra/vnc"
	"log"
	"net"
)

func main() {
	nc, err := net.Dial("tcp", "192.168.0.100:5900")
	if err != nil {
		log.Println(err)
		return
	}

	cc1, err := vnc.Client(nc, &vnc.ClientConfig{Auth: []vnc.ClientAuth{&vnc.PasswordAuth{Password: "test"}}})
	if err != nil {
		log.Println(err)
	} else {
		cc1.Close()
	}
}
