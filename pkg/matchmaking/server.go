package matchmaking

import (
	"fmt"
	"net"
)

type GameServer struct {
}

func NewServer(serverPort int) {
	address := fmt.Sprintf("0.0.0.0:%d", serverPort)
	server, err := net.Listen("tcp", address)

}
