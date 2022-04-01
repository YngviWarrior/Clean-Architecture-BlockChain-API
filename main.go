package main

import (
	"clean-go/usecase"
	"fmt"
)

func main() {
	t, err := usecase.GetTransactions()
	fmt.Println(t, err)
}
