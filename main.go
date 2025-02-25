package main

import (
	"fmt"
	"log"
	"strings"

	//"os"

	"github.com/xuri/excelize/v2"
)


func main(){
	//file_path := os.Args[1]
	
	file, err := excelize.OpenFile("./CSF111_202425_01_GradeBook_stripped.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	sheetname := file.GetSheetName(0)

	rows, err := file.GetRows(sheetname)
	if err != nil {
		log.Fatal(err)
	}

	//Checking if any of the cell in any row is blank
	for row_index, row := range rows {
		var flag = false
		for _, cell := range row {
			if strings.TrimSpace(cell) == "" {
				flag = true
				break
			}
		}
		if !flag {
			fmt.Printf("Row:%d, %v\n", row_index+1,row)
		}
	}
}