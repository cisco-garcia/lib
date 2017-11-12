package notification

import (
	"github.com/0xAX/notificator"
)

func DisplayNotification(title string, notification string) {
	notify := notificator.New(notificator.Options{
		AppName: "Notifier",
	})
	notify.Push(title, notification, "", notificator.UR_CRITICAL)
}
