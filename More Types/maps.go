package main

import "fmt"

type Person struct{
	name string
	age int
	occupation string
}

type Hero struct{
	name string
	alias string
	villan string
}

var Heros = map[string]Hero{
	"Batman":Hero{
		"Bruce Wayne","Batman","Joker",
	},
	"Superman":Hero{
		"Clark Kent","Superman","Lex Luthor",
	},
}

//a map maps keys to a value
//The key is a string and the value is of type Vertex
var family map[string]Person

func main() {
	family = make(map[string]Person)
	family["Danish"] = Person{
		"Danish", 28, "Accountant",
	}
	family["Emaan"] = Person{
		"Emaan", 33, "Mother/Cook/Wedding Decorator",
	}
	fmt.Println(family["Danish"])
	fmt.Println(family["Emaan"])

	fmt.Println(Heros)
}