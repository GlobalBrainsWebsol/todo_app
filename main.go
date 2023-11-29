package main

import (
	"controllers"
	"fmt"
	"strings"
)

func main() {
	var s = strings.Contains("sss", "aa")
	fmt.Println(s)
	controllers.StartMainServer()
}
