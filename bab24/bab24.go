package main

import "fmt"

var student struct {
	name string
	age  int
}

func panggilSiswa() {
	siswa := student
	siswa.name = "jidan"
	siswa.age = 22

	fmt.Println(student)
}
