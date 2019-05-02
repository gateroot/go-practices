package main

import . "go-practices/10.observer_pattern/observer"

func main() {
	var teacher Observer = &Teacher{}
	var student Student = Student{Subject: Subject{}}
	student.AddObserver(teacher)
	student.Run(10)
	student.Run(5)
}
