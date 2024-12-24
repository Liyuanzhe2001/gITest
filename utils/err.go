package utils

import (
	"fmt"
)

var Exception exception

type exception struct{}

func (e *exception) ErrorHandle(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func (e *exception) ErrorPrint(err ...string) {
	for s := range err {
		fmt.Print(s, " ")
	}
	// 结束程序
	panic("")
}

func (e *exception) AddPrefixErrorHandle(prefix string, err error) {
	panic(prefix + err.Error())
}

func (e *exception) AddSuffixErrorHandle(suffix string, err error) {
	panic(err.Error() + suffix)
}

func (e *exception) WarnHandle(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (e *exception) WarnPrint(err ...string) {
	for s := range err {
		fmt.Print(s, " ")
	}
}

func (e *exception) AddPrefixWarnHandle(prefix string, err error) {
	fmt.Println(prefix + err.Error())
}

func (e *exception) AddSuffixWarnHandle(suffix string, err error) {
	fmt.Println(err.Error() + suffix)
}
