package readimage

import (
	"io/ioutil"
)

func readFileList(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

}
