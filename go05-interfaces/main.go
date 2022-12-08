package main

import (
	"fmt"
	"github.com/bitfield/script"
)

func main() {
	contents, _ := script.File("test.txt").String()
	fmt.Println(contents)
}
