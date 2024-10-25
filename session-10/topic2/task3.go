package topic2

import "time"

func SendingReceiving(ch chan int, value int) {
	time.Sleep(500 * time.Millisecond)
	ch <- 42

}
