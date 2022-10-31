package assetReader

// import (
// 	"fmt"

// 	"github.com/xuri/excelize/v2"
// )

// var (
// 	var SheetValue
// )

// func ExcelReader(path string){
// 	f, err := excelize.OpenFile(path)
//     if err != nil {
// 		panic(err)
//     }

//     defer func() {
//         // Close the spreadsheet.
//         if err := f.Close(); err != nil {
// 			panic(err)
//         }
//     }()

//     // Get value from cell by given worksheet name and cell reference.
//     cell, err := f.GetCellValue("시트16", "C6")
//     if err != nil {
//         fmt.Println(err)
//         return
//     }

//     fmt.Println(cell)

//     // Get all the rows in the Sheet1.
//     // rows, err := f.GetRows("시트16")
//     // if err != nil {
//     //     fmt.Println(err)
//     //     return
//     // }

//	    // for _, row := range rows {
//	    //     for _, colCell := range row {
//	    //         fmt.Print(colCell, "\t")
//	    //     }
//	    //     fmt.Println()
//	    // }
//	}
