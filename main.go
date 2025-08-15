package main

import (
	"fmt"

	"github.com/piotrusio/go-pro/algorithms"
)


func main() {
	// input strs = ["flower","flow","flight"]
	// output "fl"
	strs := []string{"flower","flowe","floweht"}
	prefix := algorithms.LongestCommonPrefix(strs)
	fmt.Println(prefix)
}