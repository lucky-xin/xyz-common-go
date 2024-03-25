package file

import (
	"os"
	"path/filepath"
)

func CreateIfNotExist(fp string) (err error) {
	if inf, err := os.Stat(fp); os.IsNotExist(err) {
		if inf.IsDir() {
			return os.Mkdir(fp, os.ModePerm)
		}
		parent := filepath.Dir(fp)
		return os.Mkdir(parent, os.ModePerm)
	}
	return nil
}
