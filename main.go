package main

import (
	"fmt"
	"strings"
)

func main() {

  // The string package is a Go standard library package that contains functions to manipulate UTF-8 encoded strings.
  str1 := "educative.io"
  str2 := "appleisappleisappl"
  str3 := "google"

  
  // The Contains() method
  // The Contains() method can be used in Go to check if a substring is contained in a string.
  fmt.Println(str1, "io", strings.Contains(str1, "io"))
	fmt.Println(str1, "shot", strings.Contains(str1, "shot"))
	fmt.Println(str1, "", strings.Contains(str1, ""))
  
  
  // The Count() method
  // The string package provides the Count() method, which can be used to count the number of times a non-overlapping 
  // instance of a substring appears in a string in Go.
  fmt.Println(str1 , "e", strings.Count(str1, "e"))
	fmt.Println(str2, "apple", strings.Count(str2, "apple"))
  fmt.Println(str3, "", strings.Count(str3, ""))
  
  
  // The ContainsAny() method
  // The strings package provides the ContainsAny() method, which can be used to check if the input string 
  // contains any of the characters or Unicode code points of another string (passed as another input to the function).
  fmt.Println(str1, "i", strings.ContainsAny(str1, "i"))
	fmt.Println(str1, "xf", strings.ContainsAny(str1, "xf"))
}

