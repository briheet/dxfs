package p2p

// Represents remote node
type Peer interface {
}

// Handles communication between node in the network
// form (Tcp, Udp, Websockets...)
type Transport interface {
	ListenAndAccept() error
}
