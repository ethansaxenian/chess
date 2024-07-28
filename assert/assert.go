package assert

import (
	"log"
	"log/slog"
)

var context = map[string]any{}

func Assert(condition bool, msg string) {
	if !condition {
		for k, v := range context {
			slog.Error("context", k, v)
		}
		log.Fatal(msg)
	}
}

func AddContext(key string, value any) {
	context[key] = value
}

func DeleteContext(key string) {
	delete(context, key)
}

func ErrIsNil(err error, msg string) {
	Assert(err == nil, msg)
}

func Raise(msg string) {
	Assert(false, msg)
}
