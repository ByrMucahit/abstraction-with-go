package handler

import "testing"

type mockNotifier struct {
}

func (m *mockNotifier) Notify(text string) {

}

func TestRegisterHandle(t *testing.T) {
	mockNotif := &mockNotifier{}
	handler := NewRegistrationHandler(mockNotif)
	handler.Handle()
}
