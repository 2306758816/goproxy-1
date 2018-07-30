package main

import (
	"os"
	"fmt"
)

func main(){

	if len(os.Args)==1{
		fmt.Println("usage: goproxy <proxy address>")
		return
	}
	Run(os.Args[1])
}
