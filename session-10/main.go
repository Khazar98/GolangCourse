package main

func main() {

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//// Topic1 Task 1 and Task2
	// go topic1.PrintNumbers(4, 100*time.Microsecond)
	// time.Sleep(time.Second)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//Topic 2 Task3
	// value := 42
	// mych := make(chan int)
	// go topic2.SendingReceiving(mych, value)
	// val := <-mych
	// fmt.Println("Value ", val)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//Topic 2 Task4

	// ch := make(chan int)
	// go func() {
	// 	for i := 0; i <= 5; i++ {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }()
	// for val := range ch {
	// 	fmt.Println(val)
	// }

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Topic 3 Task5
	// ch := make(chan int, 3)

	// ch <- 10
	// ch <- 20
	// ch <- 30
	// go topic3.BufferedChannelOperations(ch)
	// close(ch)

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Topic 3 Task6
	ch := make(chan string)

	ch <- "Hello"

}