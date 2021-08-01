package base

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

// Replace 替换某个文件夹下所有文件的某个字符串
// @param path 文件夹路径
// @param old 将被替换的字符串
// @param new 新的字符串
func Replace(path, old, new string) {
	files := getAllFiles(path)

	for _, file := range files {
		output, needHandle, err := newFileContent(file, old, new)
		if err != nil {
			panic(err)
		}

		if needHandle {
			err = writeToFile(file, output)
			if err != nil {
				panic(err)
			}
		}
	}
}

// getAllFiles 获取路径下的所有文件名
// @param path 路径
// @return []string
func getAllFiles(path string) []string {
	files := make([]string, 0)
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(fmt.Sprintf("filepath.Walk() returned %v\n", err))
	}
	return files
}

// newFileContent  新的文件内容
// @param filePath 旧文件路径
// @param old 将被替换的字符串
// @param new 新的字符串
// @return []byte 替换后的文件内容
// @return bool 是否有修改
// @return error
func newFileContent(filePath, old, new string) ([]byte, bool, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, false, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	needHandle := false
	output := make([]byte, 0)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return output, needHandle, nil
			}
			return nil, needHandle, err
		}

		if ok, _ := regexp.Match(old, line); ok {
			reg := regexp.MustCompile(old)
			newByte := reg.ReplaceAll(line, []byte(new))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
			if !needHandle {
				needHandle = true
			}
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}
	}

	return output, needHandle, nil
}

// writeToFile 写文件
// @param filePath 文件路径
// @param outPut
// @return error
func writeToFile(filePath string, outPut []byte) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(f)
	_, err = writer.Write(outPut)
	if err != nil {
		return err
	}

	writer.Flush()
	return nil
}
