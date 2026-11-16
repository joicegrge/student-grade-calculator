package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const numMarks = 3

// reader reads full lines from the keyboard, so names with spaces work.
var reader = bufio.NewReader(os.Stdin)

// readLine reads one full line of input. It exits cleanly if input ends.
func readLine() string {
	line, err := reader.ReadString('\n')
	if err != nil && line == "" {
		fmt.Println("\nInput ended. Exiting.")
		os.Exit(1)
	}
	return strings.TrimSpace(line)
}

// calculate returns the total and the average of the marks.
func calculate(marks []int) (int, float64) {
	total := 0
	for _, m := range marks {
		total += m
	}
	average := float64(total) / float64(len(marks))
	return total, average
}

// getGrade returns the grade letter for an average.
func getGrade(average float64) string {
	if average >= 80 {
		return "A"
	} else if average >= 60 {
		return "B"
	} else if average >= 40 {
		return "C"
	}
	return "F"
}

// readName keeps asking until a non-empty name is entered.
func readName() string {
	for {
		fmt.Print("Enter student name: ")
		name := readLine()
		if name != "" {
			return name
		}
		fmt.Println("  Name cannot be empty.")
	}
}

// readMark keeps asking until the user enters a whole number from 0 to 100.
func readMark(n int) int {
	for {
		fmt.Printf("Enter mark %d: ", n)
		mark, err := strconv.Atoi(readLine())
		if err != nil {
			fmt.Println("  Invalid input: please enter a whole number.")
			continue
		}
		if mark < 0 || mark > 100 {
			fmt.Println("  Invalid mark: must be between 0 and 100.")
			continue
		}
		return mark
	}
}

func main() {
	name := readName()

	// Loop to collect marks and store them in a slice
	marks := []int{}
	for i := 1; i <= numMarks; i++ {
		marks = append(marks, readMark(i))
	}

	total, average := calculate(marks)
	grade := getGrade(average)

	result := "Fail"
	if average >= 40 {
		result = "Pass"
	}

	fmt.Println()
	fmt.Println("Student:", name)
	fmt.Println("Total:", total)
	fmt.Printf("Average: %.2f\n", average)
	fmt.Println("Grade:", grade)
	fmt.Println("Result:", result)
}
