package main

import (
	"fmt"

	"github.com/hbl-duytv/repo-1/internal"
	"github.com/hbl-duytv/repo-1/pkg/chaining"
)

func init() {
	internal.InitDB()
}

func main() {
	fmt.Println("a")
	chaining.C()
}
