package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func getInput(input chan int) {
	for {
		var result int
		_, err := fmt.Scan(&result)
		if err != nil {
			log.Fatal(err)
		}
		input <- result
	}
}

func main() {

	var (
		fib     [40]int
		correct int
		err     int
	)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i < 40; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	fmt.Println("Input the Fibonacci sequence (begin with 0):")
	input := make(chan int, 1)
	go getInput(input)
	for i := 0; i < 40; i++ {
		fmt.Print("F", i, ": ")
		select {
		case inp := <-input:
			if inp != fib[i] {
				fmt.Println("error: ")
				index, _ := json.Marshal(i)
				number, _ := json.Marshal(fib[i])
				fmt.Println("F", string(index), ": ", string(number))
				err++
				correct = 0
			}
			if inp == fib[i] {
				correct++
			}
		case <-time.After(10 * time.Second):
			fmt.Println("Time out!")
			index, _ := json.Marshal(i)
			number, _ := json.Marshal(fib[i])
			fmt.Println("F", string(index), ": ", string(number))
			err++
			correct = 0
		}
		if err == 3 {
			fmt.Println("You have 3 wrong answer!")
			break
		}
		if correct == 10 {
			fmt.Println("You gave 10 right answer! Good done!")
			break
		}
	}
}
