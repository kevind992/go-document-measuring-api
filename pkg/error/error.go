package error

import "log"

func Check(msg string, err error) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
