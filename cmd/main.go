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
	r       reader.Reader
	Version string
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please input the filename")
		os.Exit(1)
	}

	if len(os.Args) == 2 && os.Args[1] == "-v" {
		lib.DisplayVersion(Version)
		os.Exit(0)
	}

	ext := strings.ToUpper(filepath.Ext(os.Args[1]))

	// create reader
	switch ext {
	case ".TXT":
		r = reader.Reader(reader.NewTxtReader())
		if err := r.Load(os.Args[1]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unsupported file format!")
		os.Exit(1)
	}

	core.Init(r)
}
