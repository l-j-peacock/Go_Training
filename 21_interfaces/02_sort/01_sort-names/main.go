package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
}

type people interface {

}

func (p person) getName() string {
	return p.name
}

func main() {

}