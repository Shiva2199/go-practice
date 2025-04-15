package main

import "fmt"

// inheritance with composition and method overriding

type Animal struct {
	Name string
	Age  int
}

func (a Animal) Speak() string {
	return "Animal speaks"
}

func (a Animal) GetNameAndAge() string {
	return a.Name + " is " + fmt.Sprint(a.Age) + " years old"
}

type Dog struct {
	Animal
}

func (d Dog) Speak() string {
	return "Dog barks"
}

func main() {
	d := Dog{
		Animal: Animal{
			Name: "Dog",
			Age:  5,
		},
	}
	fmt.Println(d.Speak() + " and " + d.GetNameAndAge())

}
