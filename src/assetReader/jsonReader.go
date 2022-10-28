package assetReader

import (
	"encoding/json"
	"os"
	"strings"
)

var (
	Random [][]string
	RandomStringValue RandomString
)

type RandomString struct {
	Value []string
}

func JsonReader(path string) {
	jsonFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonFile, &RandomStringValue)

	Random =  make([][]string, 10500)
	for i, j := range RandomStringValue.Value {
		Random[i] = make([]string, 7)
		Random[i] = strings.Split(j, "-")
	}
}

// func JsonReader(path string){
// 	jsonInfo, err := os.ReadDir(path)
// 	if err != nil {
// 		panic(err)
// 	}

// 	JsonInfo = make([]fs.DirEntry, len(jsonInfo))
// 	Jsons = make([]HPJson, len(jsonInfo))

// 	var wg sync.WaitGroup
// 	wg.Add(len(jsonInfo))
// 	for i := range jsonInfo {
// 		go func (idx int){
// 			defer wg.Done()


// 		}(i)
// 	}
// 	wg.Wait()
// }
