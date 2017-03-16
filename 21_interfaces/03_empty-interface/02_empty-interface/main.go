package main

import "fmt"

type vehicles interface {}

type vehicle struct {
	Seats 		int
	MaxSpeed 	int
	Colour 		string
}

type car struct {
	vehicle
	Wheels	int
	Doors 	int
}

type plane struct {
	vehicle
	Jet	bool
}

type boat struct {
	vehicle
	Length 	int
}

func (v vehicle) Specs() {
	fmt.Printf("Seats %v, max speed %v, colour %v", v.Seats, v.MaxSpeed, v.Colour)
}

func main() {
	prius := car{}
	tacoma := car{}
	bmw528 := car{}

	boeing747 := plane{}
	boeing757 := plane{}
	boeing767 := plane{}

	sanger := boat{}
	nautique := boat{}
	malibu := boat{}

	rides := []vehicles{prius, tacoma, bmw528, boeing747, boeing757, boeing767, sanger, nautique, malibu}

	for key, value := range rides{
		fmt.Println(key, " - ", value)
	}
}