package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sort"
	//"sync"

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

func general_averages(students []Student) {
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

func getBranch(campusID string) string {
	if len(campusID) < 6 {
		return "Unknown"
	}
	return campusID[4:6]
}

func branch_averages(students []Student) {
	branchSums := make(map[string]float64)
	branchCounts := make(map[string]int)

	for _, s := range students {
		branch := getBranch(s.CampusID)
		branchSums[branch] += s.Total
		branchCounts[branch]++
	}

	fmt.Println("\n ||Branch-wise Averages||")
	for branch, sumTotal := range branchSums {
		count := float64(branchCounts[branch])
		avgTotal := sumTotal / count
		fmt.Printf("Branch %s: Average Total = %.2f (based on %d students)\n", branch, avgTotal, branchCounts[branch])
	}
}



func getter(s Student, category string) float64 {
	switch category {
	case "Quiz":
		return s.Quiz
	case "MidSem":
		return s.MidSem
	case "LabTest":
		return s.LabTest
	case "WeeklyLabs":
		return s.WeeklyLabs
	case "PreCompre":
		return s.PreCompre
	case "Compre":
		return s.Compre
	case "Total":
		return s.Total
	}
	return 0
}

func rankTop3(students []Student) {
	scoreCategories := map[string][]Student{
		"Quiz":        students,
		"MidSem":      students,
		"LabTest":     students,
		"WeeklyLabs":  students,
		"PreCompre":   students,
		"Compre":      students,
		"Total":       students,
	}

	fmt.Println("\n || Top 3 Students for Each Component ||")
	for category, list := range scoreCategories {
		
		sort.SliceStable(list, func(i, j int) bool {
			switch category {
			case "Quiz":
				return list[i].Quiz > list[j].Quiz
			case "MidSem":
				return list[i].MidSem > list[j].MidSem
			case "LabTest":
				return list[i].LabTest > list[j].LabTest
			case "WeeklyLabs":
				return list[i].WeeklyLabs > list[j].WeeklyLabs
			case "PreCompre":
				return list[i].PreCompre > list[j].PreCompre
			case "Compre":
				return list[i].Compre > list[j].Compre
			case "Total":
				return list[i].Total > list[j].Total
			}
			return false
		})

		fmt.Printf("\n ||%s Ranking||\n", category)
		for i, student := range list[:min(3, len(list))] {
			fmt.Printf("%d. EmplID: %s | Marks: %.2f\n", i+1, student.EmplID, getter(student, category))
		}
	}
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
			CampusID: strings.TrimSpace(row[3]),
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
	general_averages(students)
	branch_averages(students)
	rankTop3(students)
}