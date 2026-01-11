package main

import (
	"github.com/coerschkes/hyprland-event-daemon/src/api"
	"github.com/coerschkes/hyprland-event-daemon/src/api/handlers"
	"github.com/coerschkes/hyprland-event-daemon/src/hyprland"
	event_handlers "github.com/coerschkes/hyprland-event-daemon/src/hyprland/event-handlers"
	"github.com/coerschkes/hyprland-event-daemon/src/state"
)

func main() {
	pwdRegistry := state.NewPwdRegistry()

	hyprlandEventObserver := hyprland.EventObserver{
		EventHandlers: []hyprland.EventHandler{
			event_handlers.NewFocusedWindowHandler(pwdRegistry),
		},
	}

	socketServer := api.SocketServer{
		MessageHandlers: []api.MessageHandler{
			handlers.NewPwdRetrieveHandler(pwdRegistry),
		},
	}

	go hyprlandEventObserver.Start()
	socketServer.Start()
}
