package readimage

import (
	"image"
	"sync"
)

func ReadImage(paths []string) {
	var wg sync.WaitGroup
	wg.Add(len(paths))

	images := make([]image.Image, len(paths))

	for i, path := range paths {
		go func(idx int, filePath string) {
			defer wg.Done()

			var err error

			images[idx], err = readPngImage(filePath)
			if err != nil {
				panic(err)
			}
		}(i, path)
	}
	wg.Wait()
}
