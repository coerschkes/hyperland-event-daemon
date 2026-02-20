package hyprland

import "fmt"

type StartupExecutor struct {
	StartupHandlers []StartupHandler
}

func (s *StartupExecutor) Start() {
	for _, handler := range s.StartupHandlers {
		err := handler.OnStartup()
		if err != nil {
			fmt.Println("Startup error: ", err)
		}
	}
}
