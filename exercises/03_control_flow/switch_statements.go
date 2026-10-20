package main

import "fmt"

func main() {
	day := 3
	grade := 'B'
	score := 85
	
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day")
	}
	
	switch grade {
	case 'A', 'a':
		fmt.Println("Excellence")
	case 'B', 'b':
		fmt.Println("Good")
	case 'C', 'c':
		fmt.Println("Average")
	case 'D', 'd', 'F', 'f':
		fmt.Println("Needs improvement")
	default:
		fmt.Println("Invalid grade")
	}
	
	// TODO: Write a switch statement with no expression (works like if-else if)
	// Check score ranges:
	// score >= 90: "Outstanding"
	// score >= 80: "Very Good"  
	// score >= 70: "Good"
	// score >= 60: "Satisfactory"
	// default: "Needs work"
	switch {
	case score >= 90:
		fmt.Println("Outstanding")
	case score >= 80:
		fmt.Println("Very Good")
	case score >= 70:
		fmt.Println("Good")
	case score >= 60:
		fmt.Println("Satisfactory")
	default:
		fmt.Println("Needs Work")
	}
	
	// TODO: Write a switch with fallthrough
	// For day value:
	// case 1: print "Start of work week" and fallthrough to next case
	// case 2, 3, 4: print "Weekday"
	// case 5: print "TGIF!"
	// case 6, 7: print "Weekend"
	
	// TODO: Write a switch with initialization
	// Initialize x := score / 10 in the switch
	// Use x for cases: 10,9: "A", 8: "B", 7: "C", 6: "D", default: "F"
}