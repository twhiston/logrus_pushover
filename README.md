# Pushover hook for logrus [![GoDoc](http://godoc.org/github.com/toorop/logrus_pushover?status.svg)](http://godoc.org/github.com/toorop/logrus_pushover) [![Go Report Card](https://goreportcard.com/badge/github.com/toorop/logrus_pushover)](https://goreportcard.com/report/github.com/toorop/logrus_pushover)

Send Logrus log message using [Pushover](https://pushover.net/) on levels:

* Error
* Fatal
* Panic

## Installation

```go
go get github.com/toorop/logrus_pushover
```

## Usage

```go
import (
  "log/syslog"
  "github.com/Sirupsen/logrus"
  "github.com/toorop/logrus_pushover"
)

func main() {
hook, err := NewPushoverHook("PUSH_OVER_USER_TOKEN","PUSH_OVER_API_TOKEN")
	if err != nil {
		panic(err)
	}
	msg := "test message"
	log := logrus.New()
	log.Out = ioutil.Discard
	log.Hooks.Add(hook)
	log.WithFields(logrus.Fields{"fied1": "1", "field2": "2"}).Error(msg)
}
```
async:

```go
import (
  "log/syslog"
  "github.com/Sirupsen/logrus"
  "github.com/toorop/logrus_pushover"
)

func main() {
hook, err := NewPushoverAsyncHook("PUSH_OVER_USER_TOKEN","PUSH_OVER_API_TOKEN")
	if err != nil {
		panic(err)
	}
	msg := "test message"
	log := logrus.New()
	log.Out = ioutil.Discard
	log.Hooks.Add(hook)
	log.WithFields(logrus.Fields{"fied1": "1", "field2": "2"}).Error(msg)
}
