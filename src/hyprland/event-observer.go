package hyprland

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"slices"

	"github.com/coerschkes/hyprland-event-daemon/src/hyprland/domain"
)

type EventObserver struct {
	EventHandlers []EventHandler
}

func (o *EventObserver) Start() {
	socket := os.ExpandEnv("$XDG_RUNTIME_DIR/hypr/$HYPRLAND_INSTANCE_SIGNATURE/.socket2.sock")

	conn, err := net.Dial("unix", socket)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		event := domain.NewHyprlandEvent(scanner.Text())

		for _, h := range o.EventHandlers {
			if slices.Contains(h.Types(), event.Type) {
				err := h.OnEventReceived(event)
				if err != nil {
					fmt.Println("Error: ", err)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
