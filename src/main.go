package main

import "fmt"

func main() {
	//Build monkey meta records
	monkeys := GetOnChainMonkeys()

	fmt.Printf("%+v\n", monkeys[4642-1])
	fmt.Printf("%+v\n", monkeys[1179-1])
	fmt.Printf("%+v\n", monkeys[5961-1])
	fmt.Printf("%+v\n", monkeys[8059-1])
	fmt.Printf("%+v\n", monkeys[7753-1])
	fmt.Printf("%+v\n", monkeys[2301-1])
	fmt.Printf("%+v\n", monkeys[965-1])
}
