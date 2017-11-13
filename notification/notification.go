package main

import (
	"github.com/0xAX/notificator"
	"ciscogarcia.net/lib/rabbit"
)

struct Notifier {
	rabbit *Rabbit
	messages []string
}

// TODO Make notification yaml that is seperate from podcast-ripper yaml

func main() {
	// initialize rabbit
	// run message checker
	// endless loop over select() on channel
	// display messages as we get them and store them in our messages slice
}

func DisplayNotification(title string, notification string) {
	notify := notificator.New(notificator.Options{
		AppName: "Notifier",
	})
	notify.Push(title, notification, "", notificator.UR_CRITICAL)
}
