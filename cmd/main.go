package main

import (
	"fmt"
	"os"

	"github.com/TimothyYe/glance/core"
	"github.com/TimothyYe/glance/reader"
)

func main() {
	r := reader.NewTxtReader()
	if err := r.Load("/Users/timothy.ye/Downloads/dmbj.txt"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(r.Current())

	core.Init()
	core.HandleEvents()
}
