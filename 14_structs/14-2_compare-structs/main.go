package main

import "fmt"

/*
 * sctructs of same type can be compares with equatlity operators (==; !=)
 */

type s1 struct {
	x int
}

type s2 struct {
	x int
}

func main() {
	a1 := s1{x: 5}
	// b := s2{x: 5} // cannot compared to a1 because it is a different struct/type
	a2 := s1{x: 5}

	if a1 == a2 {
		fmt.Println("equal")
	} else if a1 != a2 {
		fmt.Println("different")
	} else {
		fmt.Println("Something went wrong")
	}

}
