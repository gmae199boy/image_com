package main

import (
	"flag"
	"fmt"
	"image"
	"strconv"
	"sync"
	"time"

	"image_com/src/assetReader"
	"image_com/src/combine"
)

var (
	w          = flag.Int("w", 1000, "width of image")
	h          = flag.Int("h", 1000, "heigth of image")
	imagePath  = flag.String("imagePath", "./image", "image path")
	jsonPath = flag.String("jsonPath", "./random.json", "json path")
)

func main() {
	flag.Parse()
	
	startTime := time.Now()

	assetReader.JsonReader(*jsonPath)
	
	elapsedTime := time.Since(startTime)
    fmt.Printf("json read 실행시간: %s\n", elapsedTime)

	startTime = time.Now()
	assetReader.ImageReader(*imagePath, assetReader.Images[0])
	
	elapsedTime = time.Since(startTime)
    fmt.Printf("image read 실행시간: %s\n", elapsedTime)


	// validate.Image()

	mutex := &sync.Mutex{}
	startTime = time.Now()
	var wg sync.WaitGroup
	wg.Add(2000)
	for i := 0; i < 2000; i++ {
		go func(idx int) {
			defer wg.Done()
			mutex.Lock()
			combineOpt := &combine.CombineOptions{
				W: *w,
				H: *h,
				Images: []*image.Image{
					assetReader.Images[0][assetReader.Random[idx][0]], 
					assetReader.Images[1][assetReader.Random[idx][1]],
					assetReader.Images[2][assetReader.Random[idx][2]],
					assetReader.Images[3][assetReader.Random[idx][3]],
					assetReader.Images[4][assetReader.Random[idx][4]],
					assetReader.Images[5][assetReader.Random[idx][5]],
					assetReader.Images[6][assetReader.Random[idx][6]],
				},
				OutputPath: "./result",
				SaveAs: combine.Png,
				ImageName: strconv.Itoa(idx),
			}
			mutex.Unlock()
			combine.Image(combineOpt)
		}(i)
	}
	wg.Wait()

	elapsedTime = time.Since(startTime)
    fmt.Printf("이미지 조합 실행시간: %s\n", elapsedTime)

	startTime = time.Now()
	wg.Add(2000)
	for i := 2000; i < 4000; i++ {
		go func(idx int) {
			defer wg.Done()
			mutex.Lock()
			combineOpt := &combine.CombineOptions{
				W: *w,
				H: *h,
				Images: []*image.Image{
					assetReader.Images[0][assetReader.Random[idx][0]], 
					assetReader.Images[1][assetReader.Random[idx][1]],
					assetReader.Images[2][assetReader.Random[idx][2]],
					assetReader.Images[3][assetReader.Random[idx][3]],
					assetReader.Images[4][assetReader.Random[idx][4]],
					assetReader.Images[5][assetReader.Random[idx][5]],
					assetReader.Images[6][assetReader.Random[idx][6]],
				},
				OutputPath: "./result",
				SaveAs: combine.Png,
				ImageName: strconv.Itoa(idx),
			}
			mutex.Unlock()
			combine.Image(combineOpt)
		}(i)
	}
	wg.Wait()

	elapsedTime = time.Since(startTime)
    fmt.Printf("이미지 조합 실행시간: %s\n", elapsedTime)

	startTime = time.Now()
	wg.Add(2000)
	for i := 4000; i < 6000; i++ {
		go func(idx int) {
			defer wg.Done()
			mutex.Lock()
			combineOpt := &combine.CombineOptions{
				W: *w,
				H: *h,
				Images: []*image.Image{
					assetReader.Images[0][assetReader.Random[idx][0]], 
					assetReader.Images[1][assetReader.Random[idx][1]],
					assetReader.Images[2][assetReader.Random[idx][2]],
					assetReader.Images[3][assetReader.Random[idx][3]],
					assetReader.Images[4][assetReader.Random[idx][4]],
					assetReader.Images[5][assetReader.Random[idx][5]],
					assetReader.Images[6][assetReader.Random[idx][6]],
				},
				OutputPath: "./result",
				SaveAs: combine.Png,
				ImageName: strconv.Itoa(idx),
			}
			mutex.Unlock()
			combine.Image(combineOpt)
		}(i)
	}
	wg.Wait()

	elapsedTime = time.Since(startTime)
    fmt.Printf("이미지 조합 실행시간: %s\n", elapsedTime)

	startTime = time.Now()
	wg.Add(2000)
	for i := 6000; i < 8000; i++ {
		go func(idx int) {
			defer wg.Done()
			mutex.Lock()
			combineOpt := &combine.CombineOptions{
				W: *w,
				H: *h,
				Images: []*image.Image{
					assetReader.Images[0][assetReader.Random[idx][0]], 
					assetReader.Images[1][assetReader.Random[idx][1]],
					assetReader.Images[2][assetReader.Random[idx][2]],
					assetReader.Images[3][assetReader.Random[idx][3]],
					assetReader.Images[4][assetReader.Random[idx][4]],
					assetReader.Images[5][assetReader.Random[idx][5]],
					assetReader.Images[6][assetReader.Random[idx][6]],
				},
				OutputPath: "./result",
				SaveAs: combine.Png,
				ImageName: strconv.Itoa(idx),
			}
			mutex.Unlock()
			combine.Image(combineOpt)
		}(i)
	}
	wg.Wait()

	elapsedTime = time.Since(startTime)
    fmt.Printf("이미지 조합 실행시간: %s\n", elapsedTime)

	startTime = time.Now()
	wg.Add(2500)
	for i := 8000; i < 10500; i++ {
		go func(idx int) {
			defer wg.Done()
			mutex.Lock()
			combineOpt := &combine.CombineOptions{
				W: *w,
				H: *h,
				Images: []*image.Image{
					assetReader.Images[0][assetReader.Random[idx][0]], 
					assetReader.Images[1][assetReader.Random[idx][1]],
					assetReader.Images[2][assetReader.Random[idx][2]],
					assetReader.Images[3][assetReader.Random[idx][3]],
					assetReader.Images[4][assetReader.Random[idx][4]],
					assetReader.Images[5][assetReader.Random[idx][5]],
					assetReader.Images[6][assetReader.Random[idx][6]],
				},
				OutputPath: "./result",
				SaveAs: combine.Png,
				ImageName: strconv.Itoa(idx),
			}
			mutex.Unlock()
			combine.Image(combineOpt)
		}(i)
	}
	wg.Wait()


	elapsedTime = time.Since(startTime)
    fmt.Printf("이미지 조합 실행시간: %s\n", elapsedTime)
}
