package main

import "fmt"

type Human struct {
	Name       string
	age        int
	profession string
}

func (h *Human) SayMyName() {
	fmt.Println("Say my name.")
	var declan string
	fmt.Scanln(&declan)

	if declan == "Heisenberg" {
		fmt.Println("You're goddamn right.")
	} else {
		fmt.Printf("Ok. So...\nMy name is %s\n", declan)
	}
	h.Name = declan
}

func (h Human) TalkPhone() {
	fmt.Printf("%s говорит по телефону.\n", h.Name)
}

func (h Human) Rest() {
	fmt.Printf("%s отдыхает.\n", h.Name)
}

type Action struct {
	Human
	Activity string
}

func main() {
	var personAction Action

	personAction.age = 52
	personAction.profession = "учитель химии"
	personAction.Activity = "проводит химические опыты в доме на колесах"

	personAction.SayMyName()

	fmt.Printf("Обычно %s вместе со своим другом %s.\n", personAction.Name, personAction.Activity)
	fmt.Println("Но сейчас...")
	personAction.TalkPhone()
}
