package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/TimothyYe/glance/lib"

	"github.com/TimothyYe/glance/core"
	"github.com/TimothyYe/glance/reader"
)

var (
	r       reader.GeneralReader
	Version string
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please input the file name")
		os.Exit(1)
	}

	if len(os.Args) == 2 && os.Args[1] == "-v" {
		lib.Display(Version)
		os.Exit(0)
	}

	ext := strings.ToUpper(filepath.Ext(os.Args[1]))

	// create reader
	switch ext {
	case ".TXT":
		r = reader.GeneralReader(reader.NewTxtReader())
		if err := r.Load(os.Args[1]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("File format not supported!")
		os.Exit(1)
	}

	core.Init(r)
}
