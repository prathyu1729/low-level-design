package main

import (
	"atm/atm"
	"fmt"
)

func main() {
	a := atm.NewAtm()
	if err := a.Withdraw(500); err != nil {
		fmt.Println(err.Error())
	}
}
