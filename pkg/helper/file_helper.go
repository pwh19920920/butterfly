package helper

import "os"

var FileHelper = new(fileHelper)

type fileHelper struct {
}

func (util *fileHelper) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
