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
	ch := make(chan int) // 新たなチャネルの生成
	go Generate(ch)      // Generateゴルーチンの開始
	for i := 1; i < 100000; i++ {
		prime := <-ch
		fmt.Printf("goroutines = %8d: prime number = %d\n", i, prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}
