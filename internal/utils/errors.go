package utils

import (
	"log"
)

func CheckErrors(funcName string, err error) {
	if err != nil {
		log.Fatal(funcName, err)
	}
}
func ThrowErrorsIfFalse(funcName string, true bool, err error) {
	if !true {
		log.Fatal(funcName, err)
	}
}
