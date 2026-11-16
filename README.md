# Student Grade Calculator (Go)

A command-line program written in Go. It asks for a student's name and three marks (0-100), then prints the total, average, grade and Pass/Fail result. No frontend, just the terminal.

## Features

- Asks for a student name (full names with spaces work) and 3 integer marks
- Calculates total and average using a function (`calculate`)
- Assigns a grade with `if / else` in a second function (`getGrade`)
- Prints name, total, average, grade and Pass/Fail (average >= 40)
- Validates each mark is between 0 and 100
- Uses a loop to collect the three marks
- Stores marks in a slice and sums them with `range`
- Handles invalid input gracefully (letters, out-of-range numbers and empty input ask again)

## Grading scale

| Average | Grade |
|---------|-------|
| >= 80   | A     |
| >= 60   | B     |
| >= 40   | C     |
| < 40    | F     |

A student passes if the average is 40 or more.

## Requirements

Go 1.21 or newer. Check with `go version`.

## How to run

```bash
go run main.go
```

## Example

```
Enter student name: Asha
Enter mark 1: 85
Enter mark 2: 75
Enter mark 3: 90

Student: Asha
Total: 250
Average: 83.33
Grade: A
Result: Pass
```

Invalid input example:

```
Enter mark 1: abc
  Invalid input: please enter a whole number.
Enter mark 1: 150
  Invalid mark: must be between 0 and 100.
Enter mark 1: 85
```

## How it works

- `calculate` adds the marks with `range` and returns the total (int) and average (float64). The average uses `float64(total) / float64(len(marks))` so it keeps decimals.
- `getGrade` maps the average to A, B, C or F with `if / else`.
- `readLine` reads a full line from the keyboard (so names with spaces work).
- `readMark` keeps asking until it gets a whole number from 0 to 100.
