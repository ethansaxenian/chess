package assert

import "log"

func Assert(condition bool, msg string) {
	if !condition {
		log.Fatal(msg)
	}
}
