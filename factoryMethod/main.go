package main

import "fmt"

/*
	Go doesn’t support constructors as in OO-languages,
	but constructor-like factory functions are easy to implement.
*/
type File struct {
  fd int // file descriptor number
  name string // filename
}

// the factory, which returns a pointer to the struct type
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

func main () {
	f := NewFile(10, "./test.txt")
	fmt.Println(f)

	// If File is defined as a struct type, the expressions new(File) 
	// and &File{} are equivalent. 
	fi := new(File)
	fi.fd = 10
	fi.name = "./test.txt"
	fmt.Println(fi)
}