package main

import "errors"

var (
	errInvalidNumberOfArgs = errors.New("please input only one argument")
	errInvalidPath         = errors.New("cannot open specified path")
)
