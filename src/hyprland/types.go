package hyprland

type EventHandler interface {
	Handle(event string) error
}
