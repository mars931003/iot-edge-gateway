package exceptions

import (
	"log"
)

func DeferFunc(err error, msg *string) {
	if errs := recover(); errs != nil {
	}
	defer func(errMsg *string) {
		log.Println(*errMsg)
		log.Println("recover errors")
	}(msg)
}

func ExceptionHandler(err error, msg string) {
	if err != nil {
		log.Println(msg)
	}
}
