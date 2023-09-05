package logwrapper

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

type LogWrapper struct {
	hub *sentry.Hub
}

func New(dsnURL string) LogWrapper {
	scope := sentry.NewScope()
	client, _ := sentry.NewClient(sentry.ClientOptions{
		Dsn:   dsnURL,
		Debug: true,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			log.Printf("BeforeSend event [%s]", event.EventID)
			return event
		},
	})

	hub := sentry.NewHub(client, scope)

	ConfigureLoggers()

	return LogWrapper{hub: hub}
}

func (c LogWrapper) CaptureMessage(msg string) {
	c.hub.CaptureMessage(msg)
}

func (c LogWrapper) Flush(timeout time.Duration) {
	c.hub.Flush(timeout)
}

func ConfigureLoggers() {
	logFlags := log.Ldate | log.Ltime

	sentry.Logger.SetPrefix("[sentry sdk]   ")
	sentry.Logger.SetFlags(logFlags)
	log.SetPrefix("[http example] ")
	log.SetFlags(logFlags)
}
