package main

import "errors"

var (
	ErrInvalidNumberOfArgs = errors.New("invalid number of args; please input exactly 1 argument")
	ErrEmptyKeyName        = errors.New("key name cannot be empty")
	ErrNoNewOrChangedPosts = errors.New("no new posts have been added, nor have any existing posts been changed")
)
