// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"fmt"
	"time"
)

const durationInSeconds = 100 // seconds

func main() {
	ch1 := make(chan struct{}) // HL
	ch2 := make(chan struct{}) // HL
	end := make(chan struct{}) // HL
	result := make(chan int)   // HL

	go pingPong(ch1, ch2, end, result) // HL
	go pingPong(ch2, ch1, end, result) // HL

	start := time.Now()
	ch1 <- struct{}{} // HL
	<-time.Tick(time.Second * durationInSeconds)
	close(end)     // HL
	r1 := <-result // HL
	r2 := <-result // HL
	elapsed := time.Now().Sub(start)

	fmt.Printf("Elapsed Time = %v\n", elapsed)
	fmt.Printf("%d per second\n", r1/durationInSeconds)
	fmt.Printf("%d per second\n", r2/durationInSeconds)
}

// START OMIT
func pingPong(in <-chan struct{}, // 受信
	out chan<- struct{}, // 送信
	end <-chan struct{}, // 終了指示
	result chan<- int) { // 結果
	for i := 0; ; i++ {
		select { // HL
		case v := <-in: // HL
			select { // HL
			case out <- v: // HL
			case <-end: // HL
				result <- i
				return
			}
		case <-end: // HL
			result <- i
			return
		}
	}
}

// END OMIT

// Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz, Linux(Ubuntu 15.0)
// Elapsed Time = 1m40.000249347s
// 542444 per second
// 542444 per second

// Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz, Linux(Ubuntu 15.0)
// 2016.06.28 Tip version of Go
// Elapsed Time = 1m40.000107401s
// 580596 per second
// 580596 per second

// 1.3GHz Intel Core M, MacOS X
// Elapsed Time = 1m40.000289819s
// 555193 per second
// 555193 per second
