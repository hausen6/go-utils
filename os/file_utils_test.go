package fileutils

import (
	"testing"

	"github.com/kardianos/osext"
)

func TestIsExist(t *testing.T) {
	testCase := map[string]bool{
		"file_utils.go":       true,
		"file_utils_test.go":  true,
		"not_exist_file.txt":  false,
		"not_exist_file2.txt": false,
		"../os":               true, // directory
	}

	for filename, except := range testCase {
		if isExist := IsExists(filename); isExist != except {
			t.Errorf("%v is except %v, but IsExists() return %v", filename, except, isExist)
		}
	}
}

func TestIsDir(t *testing.T) {
	testCase := map[string]bool{
		"file_utils.go":       false, // file
		"file_utils_test.go":  false, // file
		"not_exist_file.txt":  false,
		"not_exist_file2.txt": false,
		"../os":               true, // directory
	}

	for path, except := range testCase {
		if result := IsDir(path); result != except {
			t.Errorf("%v excepts %v, but result is %v", path, except, result)
		}
	}
}

func TestIsExecutable(t *testing.T) {
	curExt, _ := osext.Executable()
	testCase := map[string]bool{
		"notExitCmd":    false,
		"executable.py": true,
		curExt:          true,
	}
	for cmd, except := range testCase {
		if result := IsExecutable(cmd); result != except {
			t.Errorf("%v command excepts %v, but IsExecutable returns %v", cmd, except, result)
		}
	}
}
