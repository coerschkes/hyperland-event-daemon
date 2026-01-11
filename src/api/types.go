package api

type MessageHandler interface {
	CanHandle(msg string) bool
	Handle(msg string) (string, error)
}
