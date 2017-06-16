// A concurrent prime sieve

package main

import "fmt"

// START OMIT
// チャネルchへ2、3、4、...と順に数値を送る
func Generate(ch chan<- int) { // HL
	for i := 2; ; i++ {
		ch <- i // iをチャネルchへ送る // HL
	}
}

// チャネルinから値を読み出してチャネルoutへコピーするが、
// primeで割れる値は除く。
func Filter(in <-chan int, out chan<- int, prime int) { // HL
	for {
		i := <-in // inから値を受信 // HL
		if i%prime != 0 {
			out <- i // iをoutへ送信 // HL
		}
	}
}

// END OMIT

// 素数のふるい：連結フィルター処理
func main() {
	in := make(chan int) // 新たなチャネルの生成
	go Generate(in)      // Generateゴルーチンの開始
	for i := 1; i <= 20000; i++ {
		prime := <-in
		fmt.Printf("goroutines = %5d: prime number = %d\n", i, prime)
		out := make(chan int)
		go Filter(in, out, prime)
		in = out

		// 1回目：go Filter(in, out, 2)
		// 2回目：go Filter(in, out, 3)
		// 3回目：go Filter(in, out, 5)
		// 4回目：go Filter(in, out, 7)
		// 5回目：go Filter(in, out, 11)
		// 6回目：go Filter(in, out, 13)
		// 7回目：go Filter(in, out, 17)
		// ...
	}
}
