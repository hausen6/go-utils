package osutils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/mattn/go-colorable"
	"github.com/mitchellh/colorstring"
)

var (
	debug = log.New(colorable.NewColorableStdout(), colorstring.Color("[green]"), log.Lshortfile)
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
	return (info.Mode() & 0111) != 0
}

// FindExecutable is search out executable file in the PATH.
func FindExecutable(filename string) string {
	envpath := os.Getenv("PATH")
	pathlist := filepath.SplitList(envpath)
	for _, path := range pathlist {
		file := filepath.Join(path, filename)
		if IsExecutable(file) {
			return file
		}
	}
	return ""
}
