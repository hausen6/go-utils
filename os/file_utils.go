package fileutils

import (
	"log"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mitchellh/colorstring"
)

var (
	debug = log.New(colorable.NewColorableStdout(), "", log.Lshortfile)
)

// IsExists check the path exists or not.
func IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsDir check the path is a directory or not.
func IsDir(path string) bool {
	return false
}

// IsExecutable check the file is executable file or not.
func IsExecutable(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		// finename is not found
		return false
	}
	colorstring.Fprintf(colorable.NewColorableStdout(), "[green]DEBUG: [reset]%v: %v\n", filename, info.Mode())
	return (info.Mode() & 0111) != 0
}
