package main

import (
    "fmt"
    "regexp"
)

func main() {
    re := regexp.MustCompile(`^[A-Za-z0-9_]*@$`)
    
    // Cách 1: Dùng MatchString (tiện hơn với string)
    str := "Hello"
    result1 := re.MatchString(str)        // true
    fmt.Println(result1)
    
    // Cách 2: Dùng Match (phải convert)
    bytes := []byte(str)
    result2 := re.Match(bytes)            // true
    fmt.Println(result2)
    
}