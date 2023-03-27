package websocket

type Pool struct {
	Register   chan *Client
	UnRegister chan *Client
	Clients    map[*Clients]bool
	Broadcast  chan Message
}

NewPool()