package main

import (
	// "fmt"
	"fmt"
	// f "fmt"	// this is package alias

	"github.com/mwrks/go-basic/visibility/library"
	// . "github.com/mwrks/go-basic/visibility/library"	// using the dot will made the other packages seen as the same as main package
)

func main() {
	library.SayHello("ethan")
	// library.introduce("ethan") // this line will throws an error since it has unexported identifier

	// s1 := library.student{"ethan", 21}	// won't work since it uses the unexported struct
	// s1 := library.Student{"ethan", 21}	// won't work since it uses the unexported struct field
	s1 := library.StudentExp{"ethan", 21} // will work since it uses the exported struct
	fmt.Println("name ", s1.Name)
	// fmt.Println("grade", s1.grade)	// won't work since it uses the unexported struct field
	fmt.Println("grade", s1.Grade) // will work since it uses the exported struct field

	// sayHello("ethan")	// this function is still able to be used eventhough it's unexported because it has the same package as main
	// though when running it, it also has to be written on the go run

	fmt.Printf("Name : %s\n", library.InitStudent.Name)
	fmt.Printf("Grade : %d\n", library.InitStudent.Grade)
}
