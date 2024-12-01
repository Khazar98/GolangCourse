package topic3

import "fmt"
func BufferedChannelOperations(ch chan int ){
	for val:= range ch{
		fmt.Println(val)
	}

}