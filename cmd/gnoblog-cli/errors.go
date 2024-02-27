package main

import "errors"

var (
	ErrInvalidNumberOfArgs = errors.New("please input only one argument")
	ErrEmptyKeyName        = errors.New("key name cannot be empty")
)
