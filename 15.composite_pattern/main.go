package main

import . "composite/composite"

func main() {
	file1 := NewFile("file 1")
	file2 := NewFile("file 2")
	file3 := NewFile("file 3")
	file4 := NewFile("file 4")
	dir1 := NewDirectory("dir1")
	dir1.Add(file1)
	dir2 := NewDirectory("dir2")
	dir2.Add(file2)
	dir2.Add(file3)
	dir1.Add(dir2)
	dir1.Add(file4)

	dir1.Remove()
}