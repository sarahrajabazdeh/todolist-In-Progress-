package todolisterrors

import "log"

func ThrowError(err error) {
	log.Panic(err)
}
