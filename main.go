package main

import (
	"github.com/coerschkes/hyprland-event-daemon/src/api"
	"github.com/coerschkes/hyprland-event-daemon/src/api/api_handlers"
	"github.com/coerschkes/hyprland-event-daemon/src/hyprland"
	"github.com/coerschkes/hyprland-event-daemon/src/hyprland/handlers"
	"github.com/coerschkes/hyprland-event-daemon/src/state"
)

func main() {
	pwdRegistry := state.NewPwdRegistry()
	hyprctl := hyprland.NewHyprctl()

	unknownHandler := handlers.NewUnknownHandler()
	focusedWindowHandler := handlers.NewFocusedWindowHandler(pwdRegistry)
	mouseHandler := handlers.NewMouseHandler(hyprctl)
	monitorHandler := handlers.NewMonitorHandler(hyprctl)
	debugHandler := handlers.NewDebugHandler()

	hyprlandEventObserver := hyprland.EventObserver{
		EventHandlers: []hyprland.EventHandler{
			unknownHandler,
			focusedWindowHandler,
			mouseHandler,
			monitorHandler,
			debugHandler,
		},
	}

	hyprlandStartupExecutor := hyprland.StartupExecutor{
		StartupHandlers: []hyprland.StartupHandler{
			mouseHandler,
			monitorHandler,
		},
	}

	socketServer := api.SocketServer{
		MessageHandlers: []api.MessageHandler{
			api_handlers.NewPwdRetrieveHandler(pwdRegistry),
		},
	}

	go hyprlandStartupExecutor.Start()
	go hyprlandEventObserver.Start()
	socketServer.Start()
}
