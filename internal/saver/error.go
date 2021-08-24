package saver

import "errors"

var (
	ErrFullBuffer = errors.New("write buffer is full")
)
