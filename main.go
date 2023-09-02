package main

import (
	"github.com/GoesToEleven/puppy"
	"fmt"
)

func main () {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	s3 := puppy.BigBarks()
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}