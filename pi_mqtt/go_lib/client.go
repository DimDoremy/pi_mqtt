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
func High(address *C.char, pin uint8) {
	fmt.Println(C.GoString(address))
	conn, err := jsonrpc.Dial("tcp", C.GoString(address))
	if err != nil {
		log.Println(err)
	}

	req := RpioRequest{pin, true}
	var res RpioRequest
	err = conn.Call("RpioRequest.SetMode", req, &res)
	if err != nil {
		log.Println(err)
	}
}

//export Low
func Low(address *C.char, pin uint8) {
	fmt.Println(C.GoString(address))
	conn, err := jsonrpc.Dial("tcp", C.GoString(address))
	if err != nil {
		log.Println(err)
	}

	req := RpioRequest{pin, false}
	var res RpioRequest
	err = conn.Call("RpioRequest.SetMode", req, &res)
	if err != nil {
		log.Println(err)
	}
}

//export ReadPin
func ReadPin(address *C.char, pin uint8) bool {
	fmt.Println(C.GoString(address))
	conn, err := jsonrpc.Dial("tcp", C.GoString(address))
	if err != nil {
		log.Println(err)
	}

	req := RpioRequest{pin, false}
	var res RpioRequest
	err = conn.Call("RpioRequest.ReadPin", req, &res)
	if err != nil {
		log.Println(err)
	}
	return res.Mode
}

//export PullUp
func PullUp(address *C.char, pin uint8) {
	fmt.Println(C.GoString(address))
	conn, err := jsonrpc.Dial("tcp", C.GoString(address))
	if err != nil {
		log.Println(err)
	}

	req := RpioRequest{pin, true}
	var res RpioRequest
	err = conn.Call("RpioRequest.PullInput", req, &res)
	if err != nil {
		log.Println(err)
	}
}

//export PullDown
func PullDown(address *C.char, pin uint8) {
	fmt.Println(C.GoString(address))
	conn, err := jsonrpc.Dial("tcp", C.GoString(address))
	if err != nil {
		log.Println(err)
	}

	req := RpioRequest{pin, false}
	var res RpioRequest
	err = conn.Call("RpioRequest.PullInput", req, &res)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	fmt.Println("hello,world")
	High(C.CString("192.168.137.47:8096"), 10)
	<-time.After(time.Second * 4)
	Low(C.CString("192.168.137.47:8096"), 10)
}
