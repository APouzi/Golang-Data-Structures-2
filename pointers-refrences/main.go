package main

import "fmt"

type Class struct {
	name     string
	students []Student
}

type Student struct {
	name    string
	age     int
	classes *[]string
}

func (s *Student) changingInstnace(x int) {
	s.age = x
	fmt.Println(s.age)
}

func passFunctionHere(number *int) {
	*number = *number + 2

}

func (c *Class) changedInheritedStruc(stu *Student) {
	c.students = append(c.students, *stu)
}

func main() {

	//We create a new variable of x with value of 3, then we pass the address of x to the function and then we change that very value in main, in another function.
	x := 3
	//Here we pass the address with "&"
	passFunctionHere(&x)
	fmt.Println(x)

	classes := []string{"hello", "bye", "you"}

	alex := Student{"alex", 27, &classes}
	soha := Student{"soha", 21, &classes}
	class := &Class{"class1", []Student{}}

	p := *alex.classes
	p[0] = "changed"
	fmt.Println(classes)
	// class.students[0].name = "Hello"
	// fmt.Println(class.students[0].name)

	class.students = append(class.students, soha, alex)
	fmt.Println("some derefrenced struct ordeal", *class.students[0].classes)

	alex.changingInstnace(30)
	fmt.Println(alex.age)
}
