package main

import (
	"fmt"

	"github.com/SERV4BIZ/gfp/files"
	"github.com/SERV4BIZ/minifiers"
)

func main() {
	b, _ := files.ReadFile("/Users/watch99/GitHub/SERV4BIZ/compojs/libs/compojs.utility.js")
	buffer, _ := minifiers.JS(string(b))
	fmt.Println(buffer)
}
