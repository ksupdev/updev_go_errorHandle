package main

import (
    "fmt"
    "log"
	"github.com/ksupdev/updev_go_errorHandle/greetings"
)

func main()  {
	message, err := greetings.Hello("");
	if err != nil {
        log.Fatal(err)
    }
	fmt.Println(message)
}