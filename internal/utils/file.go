package utils

import (
	"io"
	"os"
)

func Copy(dstPath string, srcPath string) error {
	out, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer out.Close()

	in, err := os.OpenFile(dstPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer in.Close()

	_, err = io.Copy(in, out)
	return err
}
