//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(phrase string) (Event, error) {
	//In Wire, initializers are known as "providers," functions which provide a particular type.
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
