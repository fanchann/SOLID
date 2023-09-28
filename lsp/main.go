package main

import "fmt"

type Animal interface {
	Voices() string
}

type Cat struct{}

func (c Cat) Voices() string {
	return "meow meow"
}

type Dog struct{}

func (d Dog) Voices() string {
	return "woof woof"
}

type Gecko struct{}

func (g Gecko) Voices() string {
	return "ekkekekekekekek"
}

func HearAllAnimalSound(voices []Animal) map[string]interface{} {
	result := make(map[string]interface{})
	for _, voice := range voices {
		animalType := fmt.Sprintf("%T", voice)
		result[animalType] = voice.Voices()
	}
	return result
}

func main() {
	neko := Cat{}
	inu := Dog{}
	tekek := Gecko{}

	animals := []Animal{neko, inu, tekek}

	voices := HearAllAnimalSound(animals)
	fmt.Println(voices)
}
