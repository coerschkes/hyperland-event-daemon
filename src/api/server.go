package api

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const SOCKET_PATH = "/tmp/hyperland-daemon.sock"

type SocketServer struct {
	MessageHandlers []MessageHandler
}

func (s *SocketServer) Start() {
	os.Remove(SOCKET_PATH)

	listener, err := net.Listen("unix", SOCKET_PATH)

	if err != nil {
		panic(err)
	}

	defer os.Remove(SOCKET_PATH)
	defer listener.Close()

	fmt.Println("Listening on socket", SOCKET_PATH)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error during listener accept: ", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *SocketServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.TrimSpace(msg)

		isHandled := false
		for _, h := range s.MessageHandlers {
			if h.CanHandle(msg) {
				isHandled = true
				result, err := h.Handle(msg)
				if err != nil {
					fmt.Println(err.Error())
					writer.WriteString(err.Error())
					writer.Flush()
					break
				}
				if result != "" {
					fmt.Println("RESULT: " + result)
					writer.WriteString(result)
					writer.Flush()
				}
				break
			}
		}

		if !isHandled {
			fmt.Println("No handler found to handle msg: " + msg)
		}
	}
}
