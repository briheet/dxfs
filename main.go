package main

import (
	"log"

	"github.com/briheet/dxfs/p2p"
)

func main() {

	tr := p2p.NewTCPTransport(":3000")
	err := tr.ListenAndAccept()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	select {}
}
