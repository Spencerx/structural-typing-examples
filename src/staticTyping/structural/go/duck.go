package main

import "fmt"

type Duck interface {
	quack() string
	dance() string
}

type Wolf struct {
}

type Mallard struct {
}


func (mallard Mallard) quack() string { return "quack quack" }
func (mallard Mallard) dance() string  { return " _/¯ " }

func (wolf Wolf) quack() string { return "QUACK QUACK WHOO" }
func (wolf Wolf) dance() string  { return " ¯\\_()_/¯ " }
func (wolf Wolf) eat(duck Duck) string  { return " 😈 " }


func main() {
	var wolf Wolf
	var theDuck Duck

	theDuck = &wolf /* Type inference between Wolf and Duck. & used not to copy the wolf but to use its reference. */

	var aDuck Mallard

	twoDucksAlone(aDuck, theDuck)

	fmt.Printf("%s\n", wolf.eat(aDuck))
}

func twoDucksAlone(d1 Duck, d2 Duck) {
	fmt.Printf("%s\n", d1.quack())
	fmt.Printf("%s\n", d2.quack())

	fmt.Printf("%s\n", d1.dance())
	fmt.Printf("%s\n", d2.dance())
}
