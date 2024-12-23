package utils

import (
	"fmt"
	"gITest/enum"
)

var Err = &err{}

type err struct{}

func (e *err) Handle(type_ enum.Error, err error) {
	if type_ == enum.ERROR {
		panic(err.Error())
	} else {
		fmt.Println(err.Error())
	}
}
