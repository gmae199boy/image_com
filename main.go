package main

import (
	"flag"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	w          = flag.Int("w", 1920, "width of image")
	h          = flag.Int("h", 1080, "heigth of image")
	imagePath  = flag.String("imagePath", "./", "image path")
	definePath = flag.String("defieneString", "./", "define path")
)

func main() {
	flag.Parse()

	files, err := ioutil.ReadDir(*imagePath)
	if err != nil {
		log.Fatalln(err)
	}

	arr := make([]image.Image, len(files))

	var wg sync.WaitGroup

	wg.Add(len(files))
	for i, file := range files {
		go func(li int, imageFile fs.FileInfo) {
			defer wg.Done()
			log.Println(imageFile.Name())
			fileToImage, parseErr := os.Open(*imagePath + "/" + imageFile.Name())
			if parseErr != nil {
				log.Fatalln(parseErr)
			}
			defer fileToImage.Close()
			log.Println(*fileToImage)
			arr[li], err = png.Decode(fileToImage)

		}(i, file)
	}
	wg.Wait()

	b := arr[0].Bounds()
	image3 := image.NewRGBA(b)
	draw.Draw(image3, b, arr[0], image.Point{}, draw.Src)
	draw.Draw(image3, arr[1].Bounds(), arr[1], image.Point{}, draw.Over)

	result, err := os.Create("result.jpg")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(result, image3, &jpeg.Options{Quality: jpeg.DefaultQuality})
	defer result.Close()
}
