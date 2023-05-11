package main

// Go routines
// Channels
// Select

//func someFunc(num string) {
//	fmt.Println(num)
//
//}
//
//func main() {
//	go someFunc("1")
//
//	time.Sleep(time.Second * 2)
//	fmt.Println("hello there!")
//}

func main() {
	// Unbuffer channels make go routines synchronous
	//myChannel := make(chan string)
	//anotherChannel := make(chan string)
	//
	//go func() {
	//	myChannel <- "data"
	//}()
	//
	//go func() {
	//	anotherChannel <- "anotherData"
	//}()
	//
	//select {
	//case msgFromMyChannel := <-myChannel:
	//	fmt.Println(msgFromMyChannel)
	//
	//case msgFromAnotherChannel := <-anotherChannel:
	//	fmt.Println(msgFromAnotherChannel)
	//}

	//Buffer channels make fo routines async
	//charChannel := make(chan string, 3)
	//
	//chars := []string{"a", "b", "c"}
	//
	//for _, s := range chars {
	//	select {
	//	case charChannel <- s:
	//
	//	}
	//}
	//
	//close(charChannel)
	//
	//for result := range charChannel {
	//	fmt.Println(result)
	//}

}
