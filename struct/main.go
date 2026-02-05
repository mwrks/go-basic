package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type student2 struct {
	person
	age   int
	grade int
}

type student struct {
	name  string
	grade int
}

func main() {
	var s1 student
	s1.name = "john wick"
	s1.grade = 2
	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)

	// s1 := student{}
	s1.name = "wick"
	s1.grade = 2
	s2 := student{"ethan", 3}
	s3 := student{name: "jason"}
	fmt.Println("student 1 :", s1.name)
	fmt.Println("student 1 :", s1.grade)
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 2 :", s2.grade)
	fmt.Println("student 3 :", s3.name)
	fmt.Println("student 3 :", s3.grade)

	s4 := &s1
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s4.name)
	s4.name = "ethan"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s4.name)

	s5 := student2{}
	s5.name = "wick"
	s5.age = 21
	s5.person.age = 22
	s5.grade = 2
	fmt.Println("name :", s5.name)
	fmt.Println("age :", s5.age)
	fmt.Println("age :", s5.person.age)
	fmt.Println("grade :", s5.grade)

	s6 := student2{person{"Mark", 33}, 35, 8} // sub struct value assignment
	// s6 := student2{person: person{"Mark", 33}, age: 35,  grade: 8}	// this works
	// s6 := student2{person: person{"Mark", 33}}	// this also works
	fmt.Println(s6.name)
	fmt.Println(s6.person.name)
	fmt.Println(s6.person.age)
	fmt.Println(s6.age)

	// type sp7 = struct { // this is a type alias, it works but don't use this without proper usage
	type sp7 struct { // this is type definition, use this. This is local though, so use it properly
		name string
		age  int
	}
	s7 := struct {
		sp7
		address string
	}{} // it didn't work earlier because i forgot to add the empty curly brackets
	s7.sp7.name = "Adrian"
	s7.sp7.age = 44
	s7.address = "Tampa, Florida"
	fmt.Println(s7.name)
	fmt.Println(s7.sp7.name)
	fmt.Println(s7.sp7.age)
	fmt.Println(s7.address)

	s8 := struct {
		person
		grade int
	}{} // anonymous struct with empty property
	fmt.Println(s8)

	s9 := struct {
		person
		grade int
	}{
		person: person{"wick", 21},
		grade:  2,
	} // anonymous struct with filled properties
	fmt.Println(s9)

	allStudents := []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}
	for _, student := range allStudents {
		fmt.Println(student.name, "is", student.age, "years old")
	}

	allStudentsAnon := []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 21}, grade: 3},
	}
	for _, student := range allStudentsAnon {
		fmt.Println(student.name, "is", student.age, "years old has just earned score of:", student.grade)
	}

	var s10 struct {
		person
		grade int
	}
	fmt.Println(s10)

	var s11 = struct {
		person
		grade int
	}{person{allStudentsAnon[1].name, allStudentsAnon[1].age}, allStudentsAnon[1].grade}
	fmt.Println(s11)

	s12 := struct {
		person
		grade int
	}{person{allStudentsAnon[1].name, allStudentsAnon[1].age}, allStudentsAnon[1].grade}
	fmt.Println(s12)

	type studentNest struct {
		personNest struct {
			name string
			age  int
		}
		grade   int
		hobbies []string
	}
	fmt.Println(studentNest{})

	// type person struct { name string; age int; hobbies []string } // use semicolon to declare it horizontally
	// p1 := struct { name string; age int } { age: 22, name: "wick" }
	// p2 := struct { name string; age int } { "ethan", 23 }

	type person2 struct { // use backtick to add property tags
		name string `tag1`
		age  int    `tag2`
	}
	fmt.Println(person2{})

	type People = person2 // this is type alias
	p1 := person2{"wick", 21}
	fmt.Println(p1)
	p2 := People{"wick", 21}
	fmt.Println(p2)

	type People1 struct { // this is a type declaration of a struct
		name string
		age  int
	}
	fmt.Println(People1{})
	type People2 = struct { // this is type alias of an anonymous struct
		name string
		age  int
	}
	fmt.Println(People2{})

}
