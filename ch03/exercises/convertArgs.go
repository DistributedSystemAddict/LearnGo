package main

import (
	"fmt"
	"os"
)

// ArgStruct định nghĩa structure để lưu index và value của command line argument
type ArgStruct struct {
	Index int
	Value string
}

// ConvertArgsToStructs chuyển đổi os.Args thành slice of ArgStruct
func ConvertArgsToStructs(args []string) []ArgStruct {
	result := make([]ArgStruct, 0, len(args))
	
	for i, arg := range args {
		result = append(result, ArgStruct{
			Index: i,
			Value: arg,
		})
	}
	
	return result
}

func main() {
	// Lấy command line arguments
	args := os.Args
	
	// Convert sang slice of structures
	argStructs := ConvertArgsToStructs(args)
	
	// In kết quả
	fmt.Println("Command line arguments as structures:")
	fmt.Println("-------------------------------------")
	for _, argStruct := range argStructs {
		fmt.Printf("Index: %d, Value: %s\n", argStruct.Index, argStruct.Value)
	}
	
	fmt.Println("Index | Value")
	fmt.Println("------|------")
	for _, argStruct := range argStructs {
		fmt.Printf("%5d | %s\n", argStruct.Index, argStruct.Value)
	}
}