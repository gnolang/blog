package main

import "errors"

var (
	ErrInvalidNumberOfArgs = errors.New("please input only one argument")
	ErrInvalidPath         = errors.New("cannot open specified path")
)
