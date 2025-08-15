package main

import (
	"fmt"

	"github.com/piotrusio/go-pro/algo"
)


func main() {
	// input strs = ["flower","flow","flight"]
	// output "fl"
	strs := []string{"flower","flowe","floweht"}
	prefix := algo.LongestCommonPrefix(strs)
	fmt.Println(prefix)
}