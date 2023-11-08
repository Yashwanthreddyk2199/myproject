// go modinit to get the go.mod file and go mod tidy to get the packages
// gopath always set to the current directory
// go env to check the path and the root of the directory

package main // always starts with

import (
	"fmt" // contains all the basic functions
	//"myproject/simplecalc" // error in linking shows not in the src directory  package myproject/simplecalc is not in std (C:\Program Files\Go\src\myproject\simplecalc)
)

// structure to be declared always before main

type Author struct {
	Fname string
	Lname string
}
type Book struct {
	Title  string
	Pages  int
	Author Author
}

func main() {
	//a, b := 12, 3 // variable declaratrion
	//fmt.Println(simplecalc.Add(a, b))
	//fmt.Println(simplecalc.Sub(a, b))
	//fmt.Println(simplecalc.Mul(a, b))

	fmt.Println("Hello") // function always start with capital EX: Println, Add

	//2nd class

	// for i:=1; i<10; i++{    //for loop there is no precrement like ++i

	// 	fmt.Println("Hello")

	// }
	// another way of declaring of forloop similar to while loop
	// j:=0
	// for j<10 {
	// 	j=j+1
	// 	fmt.Println(j)
	// }
	/*infinite loop
	//for{
		//fmt.Println("hi")
	}*/

	// if else statement
	//other way if k:=5; num%2==0
	k := 5
	if k%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	//switch statement: no need to specify the break statement unlinke in C lang
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid Number")
	}

	//defer is used so that it can be done later
	defer fmt.Println("second")
	defer fmt.Println("third")
	fmt.Println("first") // output is first third second (follows LIFO)

	//Pointers is something which points it to the address of the variable
	n := 20
	var p *int      // p is a pointer
	p = &n          // p stores the address of n
	fmt.Println(p)  //*p=12 here we can directly change the value of n without directly referencing it
	fmt.Println(*p) // to print the value which is stored by p by using *

	// default value of a pointer is Nil. In go we cannot move the pointer by increment or decrement

	// Structures is a composite data type which holds data under single variable or name
	f := Author{
		Fname: "hi",
		Lname: "World",
	}

	g := Book{
		Title:  "hello",
		Pages:  10,
		Author: f,
	}

	fmt.Printf("%s is written by %s %s\n", g.Title, g.Author.Fname, g.Author.Lname)

	// Arrays to store values of similar datatype in a single Data structure

	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}

	for index, value := range arr {
		fmt.Println(index, value)
	}

	for _, value := range arr { // _ is used instead of index if not error is thrown
		fmt.Println(value)
	}

	//3rd class

	// slicing
	// different ways to create a slice
	//1. var s []int
	//2. s:=make([]int, 3,5)
	// s := make([]int, 3)
	//3. s:= []int{1,2,3}
	//4.
	s := arr[1:4] // var arr [5]int = [5]int{1,2,3,4,5}
	fmt.Println(s)
	s = append(s, 5)
	fmt.Println(s) // after appending
	// slicing example
	names := [4]string{ // example from go tour
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	b[0] = "XXX" // the update in b also reflects the changes in names i.e paul is change to xxx
	fmt.Println(a, b)
	fmt.Println(names)

	// Maps : it store data in key value pair. Every key is unique

	// a := make(map[string]int) here string values are key int are  values in a key value pair
	//a := map[string]int{
	// 	"one": 1
	// 	"two": 2
	// }
	m := map[string]float64{
		"x": 92.3,
		"y": 99.8,
		"z": 91.2,
	}

	m["w"] = 85.1 // to add new valuwe to a map
	xvalue := m["x"]
	fmt.Println(xvalue)
	m["x"] = 90 // updating value
	yvalue := m["x"]
	fmt.Println(yvalue)
	delete(m, "y") // to delete values

	for student, grade := range m { // to display the key pair values
		fmt.Println(student, grade)

	}

	if grade, ok := m["x"]; ok { // here ok is an bool value
		fmt.Println(grade, ok)
	} else {
		fmt.Println(" not found")

	}

}
