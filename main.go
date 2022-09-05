package main

import (
	"errors"
	"fmt"
	"github.com/k0kubun/pp/v3"
	"os"
	"time"
	"wire-demo/pkgBar"
	"wire-demo/pkgFoo"
)

type Message string

func NewMessage(phrase string) Message {
	return Message("Hi there! " + phrase)
}

type Greeter struct {
	Message Message // <- adding a Message field
	Grumpy  bool
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

type BigStructType struct {
	event   Event
	pkgFooA *pkgFoo.TypeA
	pkgFooB *pkgFoo.TypeB
	pkgBarC *pkgBar.TypeC
	pkgBarD *pkgBar.TypeD
}

func NewBigStructType(e Event, pkgFooA *pkgFoo.TypeA, pkgFooB *pkgFoo.TypeB, pkgBarC *pkgBar.TypeC, pkgBarD *pkgBar.TypeD) *BigStructType {
	return &BigStructType{
		event:   e,
		pkgFooA: pkgFooA,
		pkgFooB: pkgFooB,
		pkgBarC: pkgBarC,
		pkgBarD: pkgBarD,
	}
}

func (final *BigStructType) Run() {
	final.event.Start()
	pp.Println(final)

}

func main() {
	instance, err := InitializeBigStructTypeInstance("祝你健康!")
	if err != nil {
		fmt.Printf("failed to in half way: %s\n", err)
		os.Exit(2)
	}
	instance.Run()
}

//https://github.com/google/wire/blob/main/_tutorial/README.md
