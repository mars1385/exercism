package main

import (
	"fmt"

	"github.com/mars1385/exercism/tree"
)

func main() {

	fmt.Println(tree.Build([]tree.Record{
		{ID: 0, Parent: 1},
	}))
}
