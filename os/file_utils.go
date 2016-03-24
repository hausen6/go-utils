package fileutils

import (
	"log"
	"os"
	"runtime"

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

// IsDir :
//   与えたpathがディレクトリかどうかチェックする関数
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		// path is not exists
		return false
	}
	return info.IsDir()
}

// IsExecutable check the file is executable file or not.
// in the windows if the file exist, return true
func IsExecutable(filename string) bool {
	if !IsExists(filename) || IsDir(filename) {
		return false
	}

	if runtime.GOOS == "windows" {
		return true
	}

	info, _ := os.Stat(filename)
	colorstring.Fprintf(colorable.NewColorableStdout(), "[green]DEBUG: [reset]%v: %v\n", filename, info.Mode())
	return (info.Mode() & 0111) != 0
}
