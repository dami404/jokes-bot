package utils

import "log"

func CheckErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func ThrowErrorsIfFalse(true bool, err error) {
	if !true {
		log.Fatal(err)
	}
}
