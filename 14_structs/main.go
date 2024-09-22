package main

import "fmt"

/*
 * struct is a user defined datatype
 * a struture that groups together data elements
 * possible to reference a series of grouped values with a single var name
 * used to associate/group two or more data vars
 * e.g. a student has a name, grades, age, address ...
 */

/*
* declare struct
* type <structname> struct {
* <list of fields>
* }
 */
type Student struct {
	name string
	age  int
}

func main() {

	// initialize struct option 1
	fmt.Println("Init struct opt 1")
	var student1 Student
	fmt.Printf("%+v", student1)

	// initialize struct option 2
	fmt.Println("Init struct opt 2")
	student2 := new(Student)
	fmt.Printf("%+v", student2)

	// initialize struct option 3
	fmt.Println("Init struct opt 3")
	student3 := Student{
		name: "joe",
		age:  23,
	}
	fmt.Printf("%+v", student3)
	// also possible but not recommended
	student4 := Student{"john", 24}
	fmt.Printf("%+v", student4)

	// get fields of a sctruct
	student3.name = "not joe"
}
