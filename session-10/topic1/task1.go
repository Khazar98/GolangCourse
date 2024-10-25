package topic1

import (
	"fmt"
	"time"
)



func PrintNumbers(n int, delayTime time.Duration) {
	for i := 0; i < n; i++ {
		fmt.Println(i + 1)
	}
	time.Sleep(delayTime)

}
