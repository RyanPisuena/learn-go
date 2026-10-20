package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	} 

	for i := 2; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}

	i := 5
	// #region agent log
	if f, err := os.OpenFile("/Users/ryan.pisuena/learn-go/.cursor/debug-eb981e.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
		fmt.Fprintf(f, "{\"sessionId\":\"eb981e\",\"runId\":\"post-fix\",\"hypothesisId\":\"A\",\"location\":\"for_loops.go:16\",\"message\":\"i declared in main\",\"data\":{\"i\":%d},\"timestamp\":%d}\n", i, time.Now().UnixMilli())
		f.Close()
	}
	// #endregion
	for i > 1 {
		fmt.Println(i)
		i = i - 1
	}

	i = 1
	// #region agent log
	if f, err := os.OpenFile("/Users/ryan.pisuena/learn-go/.cursor/debug-eb981e.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
		fmt.Fprintf(f, "{\"sessionId\":\"eb981e\",\"runId\":\"post-fix\",\"hypothesisId\":\"A\",\"location\":\"for_loops.go:22\",\"message\":\"reused existing i with assignment\",\"data\":{\"i\":%d},\"timestamp\":%d}\n", i, time.Now().UnixMilli())
		f.Close()
	}
	// #endregion
	for{
		i++
		if i >= 3 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		if i != 5 {
			fmt.Println(i)
		} else {
			continue
		}
	}

	for i := 1; i < 3; i++{
		for j:= 1; j < 3; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
	}
}
