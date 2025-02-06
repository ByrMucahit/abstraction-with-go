package handler

type Notifier interface {
	Notify(text string)
}

type RegistrationHandler struct {
	notifier Notifier
}

func NewRegistrationHandler(notifier Notifier) *RegistrationHandler {
	return &RegistrationHandler{
		notifier: notifier,
	}
}

func (h *RegistrationHandler) Handle() {
	h.notifier.Notify("welcome")
}
