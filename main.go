package main

import (
	"clean-go/apis"
	"clean-go/usecase"
	"fmt"
)

func main() {
	t, err := usecase.GetTransactions()
	fmt.Println(t, err)

	resp := apis.CoinIO.FindTransaction
	fmt.Println(resp)
}
