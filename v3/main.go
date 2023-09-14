package main

/* go run -buildvcs=false . */

/*
 * extern void call_c_from_other_file();
 */

import "C"

import "fmt"

func main() {
    fmt.Println("Olá, Go!")
    C.call_c_from_other_file()
}

//export call_go_from_other_file
func call_c_from_other_file() {
    fmt.Println("Olá, Go!")
}


