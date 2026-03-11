/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"embed"
	"ioFs/cmd"
)

//go:embed static
var f embed.FS

func main() {
	cmd.Execute(f)
}
