package notification

import (
	"github.com/0xAX/notificator"
)

func CheckForNewNotifications(notifications_channel chan string) {

}

func DisplayNotification(title string, notification string) {
	notify := notificator.New(notificator.Options{
		AppName: "Notifier",
	})
	notify.Push(title, notification, "", notificator.UR_CRITICAL)
}
