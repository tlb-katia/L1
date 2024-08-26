package main

import (
	"fmt"
	"time"
)

func main() {
	customSleep(5)
	fmt.Println("Done")
}

func customSleep(duration int) {
	timeNow := time.Now()
	timeFinish := timeNow.Add(time.Duration(duration) * time.Second)

	for {
		newTime := time.Now()
		if newTime.After(timeFinish) {
			break
		}
	}
}
