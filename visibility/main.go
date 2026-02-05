package main

import "github.com/mwrks/go-basic/visibility/library"

func main() {
	library.SayHello()
	// library.introduce("ethan") // this line will throws an error since it has unexported identifier
}
