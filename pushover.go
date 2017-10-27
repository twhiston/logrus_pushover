package logrusPushover

import (
	"time"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/toorop/pushover"
)

// PushoverHook sends log via Pushover (https://pushover.net/)
type PushoverHook struct {
	async bool
	// to avoid flood, hook will wait muteDelay between
	// 2 msg sent to pushover
	muteDelay      time.Duration
	lastMsgSentAt  time.Time
	pushOverClient *pushover.Pushover
	levels 		   []logrus.Level
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
		async:     async,
		muteDelay: 15 * time.Minute,
		levels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			},
	}
	p.pushOverClient, err = pushover.NewPushover(pushoverAPIToken, pushoverUserToken)
	return &p, err
}

// Levels returns the available logging levels.
func (hook *PushoverHook) Levels() []logrus.Level {
	return hook.levels
}

func (hook *PushoverHook) SetLevels(levels []logrus.Level){
	hook.levels = levels
}

func (hook *PushoverHook) AddLevels(level logrus.Level) error {
	for _, l := range hook.levels {
		if l == level {
			return errors.New("Level already exists")
		}
	}
	hook.levels = append(hook.levels,level)
	return nil
}

// Fire is called when a log event is fired.
func (hook *PushoverHook) Fire(entry *logrus.Entry) error {
	if time.Since(hook.lastMsgSentAt) < hook.muteDelay {
		return nil
	}
	hook.lastMsgSentAt = time.Now()
	if hook.async {
		go hook.pushOverClient.Message(entry.Message)
		return nil
	}
	return hook.pushOverClient.Message(entry.Message)
}

// SetMuteDelay set muteDelay
func (hook *PushoverHook) SetMuteDelay(durationStr string) (err error) {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return err
	}
	hook.muteDelay = duration
	return nil
}
