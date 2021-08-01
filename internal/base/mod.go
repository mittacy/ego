package base

import (
	"golang.org/x/mod/modfile"
	"io/ioutil"
)

// ModulePath returns go module path.
func ModulePath(filename string) (string, error) {
	modBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return modfile.ModulePath(modBytes), nil
}
