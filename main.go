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

	hyprlandStartupExecutor := hyprland.StartupExecutor{
		StartupHandlers: []hyprland.StartupHandler{
			mouseHandler,
			monitorHandler,
		},
	}

	configReloadedHandler := handlers.NewConfigReloadedHandler(&hyprlandStartupExecutor)

	hyprlandEventObserver := hyprland.EventObserver{
		EventHandlers: []hyprland.EventHandler{
			focusedWindowHandler,
			mouseHandler,
			monitorHandler,
			debugHandler,
			configReloadedHandler,
			unknownHandler,
		},
	}

	socketServer := api.SocketServer{
		MessageHandlers: []api.MessageHandler{
			api_handlers.NewPwdRetrieveHandler(pwdRegistry),
		},
	}

	go hyprlandStartupExecutor.Execute()
	go hyprlandEventObserver.Start()
	socketServer.Start()
}
