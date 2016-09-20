package logrusPushover

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Sirupsen/logrus"
)

// get pushoverUserToken, pushoverAPIToken from ENV
func getTokensFromEnv() (pushoverUserToken, pushoverAPIToken string, err error) {
	pushoverUserToken = os.Getenv("PUSHOVER_USER_TOKEN")
	pushoverAPIToken = os.Getenv("PUSHOVER_API_TOKEN")
	if pushoverUserToken == "" || pushoverAPIToken == "" {
		err = errors.New("set env var PUSHOVER_API_TOKEN and PUSHOVER_USER_TOKEN")
	}
	return
}

func getNewHook() (hook *PushoverHook, err error) {
	pushoverUserToken, pushoverAPIToken, err := getTokensFromEnv()
	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
	return NewPushoverHook(pushoverUserToken, pushoverAPIToken)
}

func TestSync(t *testing.T) {
	hook, err := getNewHook()
	if err != nil {
		t.Error("expected err == nil, got", err)
	}
	msg := "test message"
	log := logrus.New()
	log.Out = ioutil.Discard
	log.Hooks.Add(hook)
	log.WithFields(logrus.Fields{"withField": "1", "filterMe": "1"}).Error(msg)
}

func TestAsync(t *testing.T) {
	hook, err := getNewHook()
	if err != nil {
		t.Error("expected err == nil, got", err)
	}
	msg := "test message"
	log := logrus.New()
	log.Out = ioutil.Discard
	log.Hooks.Add(hook)
	log.WithFields(logrus.Fields{"withField": "1", "filterMe": "1"}).Error(msg)
}

func TestSetDuration(t *testing.T) {
	hook, err := getNewHook()
	if err != nil {
		t.Error("expected err == nil, got", err)
	}
	err = hook.SetMuteDelay("blabla")
	if err == nil {
		t.Error("expected err != nil, got", err)
	}
	err = hook.SetMuteDelay("15m")
	if err != nil {
		t.Error("expected err == nil, got", err)
	}
}
