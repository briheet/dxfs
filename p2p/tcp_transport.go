package p2p

import (
	"fmt"
	"net"
	"sync"
)

// Remote node over a TCP established conn
type TCPPeer struct {
	conn net.Conn

	// if connects to someone then true
	// else if someone connects to it then false
	outbound bool
}

type TCPTransport struct {
	listenAddr string
	listener   net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddr: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {

	var err error

	t.listener, err = net.Listen("tcp", t.listenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {

	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		go t.handleConn(&conn) // reference

	}

}

func (T *TCPTransport) handleConn(conn *net.Conn) {
	fmt.Printf("new incoming conn: %+v\n", conn)
}
