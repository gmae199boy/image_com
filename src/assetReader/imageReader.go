package assetReader

import (
	"image"
	"image/png"
	"io/fs"
	"os"
	"strconv"
	"sync"

	"github.com/pkg/errors"
)

var (
	Images []map[string]*image.Image = make([]map[string]*image.Image, 7)
)

// Read() -imagePath 변수에 설정한 모든 이미지를 로딩
func ImageReader(path string, img map[string]*image.Image) {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(len(files))
	for i, file := range files {
		go func(idx int, fileinfo fs.DirEntry){
			defer wg.Done()

			if fileinfo.IsDir() {
				Images[idx] = make(map[string]*image.Image)
				ImageReader(path + "/" + fileinfo.Name(), Images[idx])
			} else {
				pngImage, err := readPng(path + "/" + fileinfo.Name())
				if err != nil {
					panic(err)
				}

				mutex.Lock()
				img[strconv.Itoa(idx)] = pngImage
				mutex.Unlock()
			}
		} (i, file)
	}
	wg.Wait()
}

// readPng() PNG 이미지를 로딩
func readPng(path string) (*image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer file.Close()

	pngImage, err := png.Decode(file)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &pngImage, nil
}
