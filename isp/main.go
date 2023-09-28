package main

import "fmt"

type Worker interface {
	Work()
}

type Eater interface {
	Eat()
}

type Human struct {
	Name string
}

func (h Human) Work() {
	fmt.Printf("%s is working\n", h.Name)
}
func (h Human) Eat() {
	fmt.Printf("%s eating\n", h.Name)
}

type Robot struct {
	Name string
}

func (r Robot) Work() {
	fmt.Printf("%s is working\n", r.Name)
}

func main() {
	farda := Human{"Farda Ayu"}
	robocop := Robot{"Robocop"}

	farda.Eat()
	farda.Work()

	robocop.Work()
}
