package main

import (
	//"fmt"
	"fmt"
	"log"
	"os"
	"strconv"
	"math"
	"github.com/xuri/excelize/v2"
)


type Student struct {
	ClassNo int
	EmplID string
    CampusID string
    Quiz float64
    MidSem float64
    LabTest float64
    WeeklyLabs float64
    PreCompre float64
    Compre float64
    Total float64
    ComputedSum float64
}

const tolerance = 0.0001

func computeAverages(students []Student) {
	var sumQuiz, sumMidSem, sumLabTest, sumWeeklyLabs, sumPreCompre, sumCompre, sumTotal float64
	numStudents := float64(len(students))

	if numStudents == 0 {
		fmt.Println("No student data available.")
		return
	}

	
	for _, s := range students {
		sumQuiz += s.Quiz
		sumMidSem += s.MidSem
		sumLabTest += s.LabTest
		sumWeeklyLabs += s.WeeklyLabs
		sumPreCompre += s.PreCompre
		sumCompre += s.Compre
		sumTotal += s.Total
	}

	avgQuiz := sumQuiz / numStudents
	avgMidSem := sumMidSem / numStudents
	avgLabTest := sumLabTest / numStudents
	avgWeeklyLabs := sumWeeklyLabs / numStudents
	avgPreCompre := sumPreCompre / numStudents
	avgCompre := sumCompre / numStudents
	avgTotal := sumTotal / numStudents

	fmt.Println("\n ||General Averages||")
	fmt.Printf("Quiz Average: %.2f\n", avgQuiz)
	fmt.Printf("Mid-Sem Average: %.2f\n", avgMidSem)
	fmt.Printf("Lab Test Average: %.2f\n", avgLabTest)
	fmt.Printf("Weekly Labs Average: %.2f\n", avgWeeklyLabs)
	fmt.Printf("Pre-Compre Average: %.2f\n", avgPreCompre)
	fmt.Printf("Compre Average: %.2f\n", avgCompre)
	fmt.Printf("Overall Total Average: %.2f\n", avgTotal)
}


func main(){
	file_path := os.Args[1]
	
	file, err := excelize.OpenFile(file_path)
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

	var students []Student
	var row_id = 1
	rows.Next()
	for rows.Next() {
		
		row, _ := rows.Columns()
		if len(row) == 0 {
			continue
		}
		
		classno, _ := strconv.Atoi(row[1])		
		quiz, _ := strconv.ParseFloat(row[4], 32)
		midSem, _ := strconv.ParseFloat(row[5], 32)
		labTest, _ := strconv.ParseFloat(row[6], 32)
		weeklyLabs, _ := strconv.ParseFloat(row[7], 32)
		preCompre, _ := strconv.ParseFloat(row[8], 32)
		compre, _ := strconv.ParseFloat(row[9], 32)
		total, _ := strconv.ParseFloat(row[10], 32)
		
		

		var computedSum float64 = quiz + midSem + labTest + weeklyLabs + compre
		
		student := Student{
			ClassNo: classno,
			EmplID: row[1],
			Quiz: quiz,
			MidSem: midSem,
			LabTest: labTest,
			WeeklyLabs: weeklyLabs,
			PreCompre: preCompre,
			Compre: compre,
			Total: total,
			ComputedSum: computedSum,
		}

		if math.Abs(computedSum-total) > tolerance {
			log.Printf("!!! Error in row %d: Computed sum %.2f does not match given total %.2f\n", row_id, computedSum, total)
		}
	
		students = append(students, student)
		row_id++

	}

	computeAverages(students)
}