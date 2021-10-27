package main

import (
	"fmt"
	"github.com/sarsembin/hw3Go/atoi"
	"github.com/sarsembin/hw3Go/fibonacci"
	"github.com/sarsembin/hw3Go/importsSort"
	"github.com/sarsembin/hw3Go/itoa"
	"github.com/sarsembin/hw3Go/reverse"
	"github.com/sarsembin/hw3Go/runeByIndex"
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

	generator := fibonacci.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(generator(), " ")
	}
	s := "みくだよ"
	index := 10
	r, err := runeByIndex.RuneByIndex(&s, &index)
	fmt.Printf("rune: %v, error: %v\n", string(r), err)
}
