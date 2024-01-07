package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("string join")
	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println("简单格式")
	fmt.Println(os.Args[1:])
}
