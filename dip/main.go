package main

import "fmt"

type Messenger interface {
	SendMessage(message string)
}

type Email struct{}

func (e Email) SendMessage(message string) {
	fmt.Printf("sending message using email : %s \n", message)
}

type SMS struct{}

func (s SMS) SendMessage(message string) {
	fmt.Printf("sending message using sms : %s \n", message)
}

type Notifier struct {
	Messenger
}

func (n Notifier) Notify(message string) {
	n.Messenger.SendMessage(message)
}

func NewNotifier(messenger Messenger) *Notifier {
	return &Notifier{Messenger: messenger}
}

func main() {
	emailMessenger := Email{}
	smsMessenger := SMS{}

	notifierEmail := NewNotifier(emailMessenger)
	notifierSMS := NewNotifier(smsMessenger)

	notifierEmail.Notify("Hello, From Email")
	notifierSMS.Notify("Hello, From SMS")

}
