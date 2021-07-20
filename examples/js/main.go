package main

import (
	"fmt"

	"github.com/SERV4BIZ/minifiers"
)

func main() {
	//b, _ := files.ReadFile("/Users/watch99/GitHub/SERV4BIZ/compojs/assets/js/lib/moment.min.js")
	buffer, _ := minifiers.JS(string(`
	return e += (t - s) / 12, 1 == s ? j(e) ? 29 : 28 : 31 - s % 7 % 2//test
	`))
	fmt.Println(buffer)
}
