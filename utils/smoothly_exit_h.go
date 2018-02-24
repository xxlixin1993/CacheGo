package utils

import "container/list"

var exitList *ExitList

type ExitInterface interface {
	// get module name
	GetModuleName() string

	// exit module function
	Stop() error
}

type ExitList struct {
	// exit list
	ll *list.List

	// exit module name
	module map[string]*list.Element
}
