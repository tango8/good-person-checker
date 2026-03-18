package main

import (
	"fmt"
	goodPersonMessage "github.com/tango8/good-person-checker/message"
)

func main() {
	message, _, _ := goodPersonMessage.GetMessage()
	fmt.Println(message)
}
