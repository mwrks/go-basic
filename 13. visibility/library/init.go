package library

import (
	"fmt"
)

// InitStudent is a package-level variable.
// It is exported (accessible from other packages) because it starts with a capital letter.
var InitStudent StudentExp

/*
// Alternative approach:
func init() {
	// This method mutates existing fields of the zero-initialized struct.
	// Less ideal if you want to ensure all other fields are reset to zero.
	InitStudent.Name = "John Wick"
	InitStudent.Grade = 2
	fmt.Println("--> library/library.go imported")
}
*/

func init() {
	// Use '=' (assignment) to update the package-level variable.
	// Avoid ':=' (short variable declaration) which would shadow it locally.
	// Using a struct literal ensures the entire object is initialized/overwritten.
	InitStudent = StudentExp{
		Name:  "John Wick",
		Grade: 2,
	}
	fmt.Println("--> library/library.go imported")
}
