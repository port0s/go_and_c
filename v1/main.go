package main

/* go run -buildvcs=false . */

//#include <stdio.h>
//void call_c_from_other_file(int a) {
//  printf("%d\n", a);
//  printf("Olá, C!\n");
//}

import "C"

import "fmt"

func main() {
    fmt.Println("Olá, go!")
    a := C.int(2)
    C.call_c_from_other_file(a)
}
