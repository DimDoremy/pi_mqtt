package main

import "C"
import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"time"
)

// rpio控制响应结构体
type RpioRequest struct {
	Pin  uint8
	Mode bool
}

//export High
func High(address string, pin uint8) {
	fmt.Println(address)
	conn, err := jsonrpc.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	req := RpioRequest{pin, true}
	var res RpioRequest
	err = conn.Call("RpioRequest.SetMode", req, &res)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

//export Low
func Low(address string, pin uint8) {
	fmt.Println(address)
	conn, err := jsonrpc.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	req := RpioRequest{pin, false}
	var res RpioRequest
	err = conn.Call("RpioRequest.SetMode", req, &res)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

//export ReadPin
func ReadPin(address string, pin uint8) bool {
	fmt.Println(address)
	conn, err := jsonrpc.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	req := RpioRequest{pin, false}
	var res RpioRequest
	err = conn.Call("RpioRequest.ReadPin", req, &res)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return res.Mode
}

//export PullUp
func PullUp(address string, pin uint8) {
	fmt.Println(address)
	conn, err := jsonrpc.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	req := RpioRequest{pin, true}
	var res RpioRequest
	err = conn.Call("RpioRequest.PullInput", req, &res)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

//export PullDown
func PullDown(address string, pin uint8) {
	fmt.Println(address)
	conn, err := jsonrpc.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	req := RpioRequest{pin, false}
	var res RpioRequest
	err = conn.Call("RpioRequest.PullInput", req, &res)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func main() {
	fmt.Println("hello,world")
	High("192.168.137.47:8096", 10)
	<-time.After(time.Second * 4)
	Low("192.168.137.47:8096", 10)
}
