package main

import (
	"fmt"
	"github.com/sarsembin/hw3Go/atoi"
	"github.com/sarsembin/hw3Go/importsSort"
	"github.com/sarsembin/hw3Go/itoa"
	"github.com/sarsembin/hw3Go/reverse"
)

func main() {
	itoaResult := itoa.Itoa(-12345678)
	fmt.Printf("value: %v, type: %T\n" , itoaResult, itoaResult)

	atoiResult, err := atoi.Atoi("-12354")
	fmt.Printf("value: %v, type: %T, err: %v\n" , atoiResult, atoiResult, err)
	atoiResult, err = atoi.Atoi("-123-54")
	fmt.Printf("value: %v, type: %T, err: %v\n" , atoiResult, atoiResult, err)

	fmt.Println(reverse.Reverse("社会"))

	importsSort.SortImports("importsSortTest.go")
}
