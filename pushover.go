package logrusPushover

import (
	"github.com/Sirupsen/logrus"
	"github.com/thorduri/pushover"
)

// PushoverHook sends log via Pushover (https://pushover.net/)
type PushoverHook struct {
	async          bool
	pushOverClient *pushover.Pushover
}

// NewPushoverHook init & returns a new PushoverHook
func NewPushoverHook(pushoverUserToken, pushoverAPIToken string) (*PushoverHook, error) {
	return newPushoverHook(pushoverUserToken, pushoverAPIToken, false)
}

// NewPushoverAsyncHook init & returns a new async PushoverHook
func NewPushoverAsyncHook(pushoverUserToken, pushoverAPIToken string) (*PushoverHook, error) {
	return newPushoverHook(pushoverUserToken, pushoverAPIToken, true)
}

// newPushoverHook init & returns a new PushoverHook
func newPushoverHook(pushoverUserToken, pushoverAPIToken string, async bool) (*PushoverHook, error) {
	var err error
	p := PushoverHook{
		async: async,
	}
	p.pushOverClient, err = pushover.NewPushover(pushoverAPIToken, pushoverUserToken)
	return &p, err
}

// Levels returns the available logging levels.
func (hook PushoverHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
}

// Fire is called when a log event is fired.
func (hook *PushoverHook) Fire(entry *logrus.Entry) error {
	if hook.async {
		go hook.pushOverClient.Message(entry.Message)
		return nil
	}
	return hook.pushOverClient.Message(entry.Message)
}
