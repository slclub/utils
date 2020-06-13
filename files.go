package utils

import (
	"io/ioutil"
	"os"
)

func IsFileExist(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	//我这里判断了如果是0也算不存在
	if fileInfo.Size() == 0 {
		return false, nil
	}

	return false, err
}

// Read all  contents from file.
func ReadAll(file string) (string, bool) {
	content, ok := ReadAllByte(file)
	return string(content), ok
}

// Read all  contents from file.
func ReadAllByte(file string) ([]byte, bool) {
	f, err := os.Open(file)
	if err != nil {
		return nil, false
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, false
	}

	return (fd), true
}
