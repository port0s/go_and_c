package main

/* go run -buildvcs=false . */

/*
 * extern void call_c_from_other_file();
 */

import "C"

import "fmt"

func main() {
    fmt.Println("Olá, go!")
    C.call_c_from_other_file()
}
