package main

import "fmt"

func main() {
	msg := "some msg"
	msgP := &msg
	fmt.Println(msgP)
	fmt.Println(*msgP)

	changeMsg(&msg, "Message 1")
	fmt.Println(msg)

	changeMsg(msgP, "Message 2")
	fmt.Println(msg)

	changeMsg(msgP, "Message 3")
	fmt.Println(*msgP)
}

func changeMsg(aPointer *string, newMsg string) {
	*aPointer = newMsg
}
