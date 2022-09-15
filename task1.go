package main

import (
	"fmt"
)

type Human struct {
	name      string
	birthDate string
}

func (h *Human) ChangeName(newName string) {
	h.name = newName
}

func (h *Human) ChangeDB(newDB string) {
	h.birthDate = newDB
}

type Action struct {
	Human        // inherit methods and fields, but methods use parent fields
	name  string // overlap with Human.name, so Action.ChangeName will work incorrectly
}

func MakeAction(name string, h Human) Action {
	return Action{
		name:  "action name",
		Human: h,
	}

}

func main() {
	h := Human{
		name:      "human name",
		birthDate: "21.12.2012",
	}
	a := MakeAction("action name", h)

	a.ChangeName("new name") // changes Human(!!!) name, not Action name
	a.ChangeDB("31.01.2000") // changes Human.birthDate, but Action inherits it

	fmt.Printf("Action.name: %+v\n", a.name)           // Action.name: action name
	fmt.Printf("Action.birthDate: %+v\n", a.birthDate) // Action.birthDate: 31.01.2000

}
