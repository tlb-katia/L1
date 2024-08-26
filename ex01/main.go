package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Activity string
	Human
}

func (a *Action) Hello() {
	fmt.Printf("My name is %s, I am %d yo and I can %s\n", a.Name, a.Age, a.Activity)
}

func main() {
	action := Action{
		Activity: "run",
		Human: Human{
			Name: "Katia",
			Age:  20,
		},
	}

	action.Hello()
}
