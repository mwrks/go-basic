package library

import "fmt"

func SayHello(name string) {	// has exported identifier, visible to the entire program (any package that imports it)
	fmt.Println("hello")
	introduce(name)	// it still has the unexported identifier but will work fine if the SayHello function is used 
}
func introduce(name string) {	//has unexported identifier, visible only within the package where it is defined
	fmt.Println("my name is:", name)
}
