// package validate

// import (
// 	"fmt"
// 	"image_com/src/assetReader"
// 	"image_com/src/types"
// 	"log"
// 	"regexp"
// 	"sync"
// 	"time"
// )

// // import (
// // 	"image_com/src/types"
// // 	"image_com/src/assetReader"
// // )

// // ValidateImage() json의 png value 값 들에 대한 이미지들이 있는지 확인
// func Image(){
// 	startTime := time.Now()

// 	var trig bool = false

// 	/*
// 		For Goroutine
// 		2.139088719s
// 	*/
// 	var wg sync.WaitGroup
// 	wg.Add(len(assetReader.Jsons))
// 	for _, json := range assetReader.Jsons {
// 		go func(file types.HPJson) {
// 			defer wg.Done()
// 			for _, attribute := range file.Attributes {
// 				if attribute.Value == "NONE" { break }
// 				r, _ := regexp.Compile(fmt.Sprintf("(?i)%s", attribute.Value))

// 				for _, image := range assetReader.ImageName {
// 					if r.MatchString(image) {
// 						trig = true
// 						break
// 					}
// 				}
// 				if !trig {
// 					log.Println(attribute.Value)
// 					panic("json에 있는 에셋이 로드한 이미지에 없습니다.")
// 				}

// 				trig = false
// 			}
// 		}(json)
// 	}
// 	wg.Wait()

// 	elapsedTime := time.Since(startTime)
//     fmt.Printf("실행시간: %s\n", elapsedTime)
// }

// /*
// 	No Goroutine
// 	12.002573782s
// */
// // func ValidateImage(){
// // 	startTime := time.Now()

// // 	var trig bool = false

// // 	for _, json := range assetReader.Jsons {
// // 		for _, attribute := range json.Attributes {
// // 			r, _ := regexp.Compile(fmt.Sprintf("(?i)%s", attribute.Value))

// // 			for _, image := range assetReader.ImageInfo {
// // 				if r.MatchString(image.Name()) {
// // 					trig = true
// // 					break
// // 				}
// // 			}
// // 			if !trig {
// // 				panic("json에 있는 에셋이 로드한 이미지에 없습니다.")
// // 			}
// // 		}
// // 	}

// // 	elapsedTime := time.Since(startTime)
// //     fmt.Printf("실행시간: %s\n", elapsedTime)
// // }
package validate