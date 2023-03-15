package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string){
	return len(name), strings.ToUpper(name)
}
func multiply(a,b int) int{
	return a * b
}
func repeatMe(words ...string){
	fmt.Println(words)


}
func main() {
	totalLength, upperName := lenAndUpper("qwe3r")
	fmt.Println(totalLength, upperName)

	repeatMe("dog", "cat", "zebra", "moon")
}
