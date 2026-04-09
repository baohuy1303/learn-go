package main

func main(){
	var channel = make(chan int) // use make keyword, chan (channel) type
	// HAVE TO USE with goroutines because it locks and wait for var change
	go processChannel(channel)
	for i := range channel{ //iterate over channel 
		println(i) // open and waiting for a value change
	}
}

func processChannel(c chan int){
	for i := 0; i < 5; i++{
		c <- i // assign with <- , and get value with <-
	}
	close(c) //close so that main stops listening from channel for value
}