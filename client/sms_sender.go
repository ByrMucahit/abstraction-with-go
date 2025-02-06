package client

import "fmt"

type XCompanySmsSender struct {
}

func (s *XCompanySmsSender) Notify(text string) {
	fmt.Println("x company sms sent: " + text)
}
