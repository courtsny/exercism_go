package main

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(visitor_name string) string
}

func SayHello(visitor_name string, g Greeter) {
	fmt.Printf("I can speak %s: %s", g.LanguageName(), g.Greet(visitor_name))
}

// German

type GermanGreeter struct {
	Language string
}

func (g GermanGreeter) LanguageName() string {
	return "German"
}

func (g GermanGreeter) Greet(visitor_name string) string {
	return "Hallo " + visitor_name + "!\n"
}

type ItalianGreeter struct {
	Language string
}

// Italian

func (i ItalianGreeter) LanguageName() string {
	return "Italian"
}

func (i ItalianGreeter) Greet(visitor_name string) string {
	return "Ciao " + visitor_name + "!\n"
}

// Portuguese

type PortugeseGreeter struct {
	Language string
}

func (p PortugeseGreeter) LanguageName() string {
	return "Portugese"
}

func (p PortugeseGreeter) Greet(visitor_name string) string {
	return "Olá " + visitor_name + "!\n"
}

func InterfacesTest() {
	var g = GermanGreeter{}
	SayHello("Dietrich", g)
	var i = ItalianGreeter{}
	SayHello("Luigi", i)
	var p = PortugeseGreeter{}
	SayHello("Jiao", p)
}
