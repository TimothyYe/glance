package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/TimothyYe/glance/core"
	"github.com/TimothyYe/glance/reader"
)

// Version is injected at build time via -ldflags "-X main.Version=...".
var Version string

// resolveVersion returns the build-injected version, or "dev" when none was set.
func resolveVersion(version string) string {
	if version == "" {
		return "dev"
	}
	return version
}

var rootCmd = &cobra.Command{
	Use:          "glance <file>",
	Short:        "A cross-platform command-line text novel reader",
	Version:      resolveVersion(Version),
	Args:         cobra.MaximumNArgs(1),
	SilenceUsage: true,
	RunE:         run,
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func run(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please input the filename")
	}

	r, err := createReader(args[0])
	if err != nil {
		return err
	}

	core.Init(r)
	return nil
}

// createReader selects a reader based on the file extension and loads the file.
func createReader(path string) (reader.Reader, error) {
	ext := strings.ToUpper(filepath.Ext(path))

	switch ext {
	case ".TXT":
		r := reader.Reader(reader.NewTxtReader())
		if err := r.Load(path); err != nil {
			return nil, err
		}
		return r, nil
	default:
		return nil, fmt.Errorf("unsupported file format")
	}
}

// Execute runs the root command, exiting non-zero on error.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
