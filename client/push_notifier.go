package client

import "fmt"

type PushNotifier struct {
}

func (s *PushNotifier) Notify(text string) {
	fmt.Println("push notification sent: " + text)
}
