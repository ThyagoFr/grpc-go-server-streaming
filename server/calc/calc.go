package calc

import (
	"fmt"
	"time"
)

// CalculatePrime - CalculatePrime
func CalculatePrime(number int64, channel chan int64) {
	var k int64 = 2
	defer close(channel)
	for number > 1 {
		if number%k == 0 {
			time.Sleep(1000 * time.Millisecond)
			channel <- k
			number = number / k
		} else {
			k++
		}
	}
	fmt.Println("End calc...")
}
