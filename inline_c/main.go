package main

//#include  <stdio.h>
//int Add(int a, int b){
//    printf("Welcome from external C function\n");
//    return a + b;
//}
import "C"
import "fmt"

func main() {
	fmt.Println(C.Add(5, 2))
}

// ref: https://stackoverflow.com/questions/70835519/using-a-cgo-shared-library-in-a-go-program
