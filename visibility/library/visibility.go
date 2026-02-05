package library

import "fmt"

func SayHello() {	// has exported identifier, visible to the entire program (any package that imports it)
	fmt.Println("hello")
}
func introduce(name string) {	//has unexported identifier, visible only within the package where it is defined
	fmt.Println("my name is:", name)
}
