package utils

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func numOfGoroutines(){
	for range time.Tick(time.Minute * 1) {
		fmt.Println("-----#Go Routine------> " + strconv.Itoa(runtime.NumGoroutine()))
	}
}

func LogGoRoutines() {
	go numOfGoroutines()
}
