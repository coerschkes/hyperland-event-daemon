package main

import (
	"github.com/coerschkes/hyprland-event-daemon/src/api"
	"github.com/coerschkes/hyprland-event-daemon/src/api/handlers"
	"github.com/coerschkes/hyprland-event-daemon/src/state"
)

func main() {
	pwdRegistry := state.NewPwdRegistry()

	socketServer := api.SocketServer{
		MessageHandlers: []api.MessageHandler{
			handlers.NewPwdRetrieveHandler(*pwdRegistry),
			handlers.NewPwdUpdateHandler(*pwdRegistry),
		},
	}

	socketServer.Start()
}
