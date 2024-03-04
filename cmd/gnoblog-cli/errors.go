package main

import "errors"

var (
	ErrInvalidNumberOfArgs = errors.New("invalid number of args; please input exactly 1 argument")
	ErrEmptyKeyName        = errors.New("key name cannot be empty")
)
