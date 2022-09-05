//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-demo/pkgBar"
	"wire-demo/pkgFoo"
)

func InitializeBigStructTypeInstance(phrase string) (*BigStructType, error) {
	//In Wire, initializers are known as "providers," functions which provide a particular type.
	wire.Build(NewMessage, NewGreeter, NewEvent, pkgFoo.PkgFooProviderSet, pkgBar.PkgBarProviderSet, NewBigStructType)
	return &BigStructType{}, nil
}
