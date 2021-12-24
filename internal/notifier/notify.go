package notifier

import "github.com/hachi-n/command-wrap-notify/internal/message"

type Notifier interface {
	Notify(message message.Message) ([]byte, error)
}
