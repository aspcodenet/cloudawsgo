package main

import (
	"fmt"

	"github.com/cheynewallace/tabby"
)

func main() {
	fmt.Println("Hej hej kolla vad coolt")
	t := tabby.New()
	t.AddHeader("NAME", "TITLE", "DEPARTMENT")
	t.AddLine("John Smith", "Developer", "Engineering")
	t.AddLine("Stefan Holmberg", "Developer", "Whatever")
	t.Print()
}
