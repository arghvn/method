package main

//Go supports methods defined on struct types.

//This area method has a receiver type of *rect.

import "fmt"

type rect struct {
	width, height int
}

//Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.

//Here we call the 2 methods defined for our struct.

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

//output :

//area:  50
//perim: 30
//area:  50
//perim: 30

//more details :

//Go automatically handles conversion between values and pointers for method calls.
// You may want to use a pointer receiver type to avoid copying on method calls
// or to allow the method to mutate the receiving struct.
