package main

import (
	"fmt"
	"os"

	"github.com/AhmetMuratAcar/Termi/tui"
)

func main() {
	fmt.Println("Program Launched")

	if _, err := tui.RunTUI(); err != nil {
		fmt.Printf("There has been an error: %v", err)
		os.Exit(1)
	}
}
