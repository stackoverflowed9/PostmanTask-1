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

	rows, err := file.Rows(sheetname)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//Checking whether any of the cell in any row is
	var row_id = 1
	rows.Next()
	for rows.Next() {
		row, _ := rows.Columns()
		var flag = false
		
		for _, cell := range row {
			if strings.TrimSpace(cell) == "" {
				flag = true
				break
			}
		}
		if !flag && len(row) != 0{
			fmt.Printf("%d: %v\n", row_id, row)
			row_id++
		}
	}
}