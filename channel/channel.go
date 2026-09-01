package main

// sending
/*func processnum(numchan chan int) {

	for num := range numchan {
		fmt.Println("processing number", num)
		time.Sleep(time.Second * 1)
	}

}*/

/*func sum(result chan int, num1 int, num2 int) {
	numresult := num1 + num2
	result <- numresult
}*/

/*func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("processing...")

}*/

/*func emailsender(emailchan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for email := range emailchan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}*/

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 5
	}()

	go func() {
		chan2 <- "hello"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1val := <-chan1:
			println("received number", chan1val)
		case chan2val := <-chan2:
			println("received message", chan2val)
		}
	}

	/*emailchan := make(chan string, 100)
	done := make(chan bool)

	go emailsender(emailchan, done)

	for i := 0; i < 10; i++ {
		emailchan <- fmt.Sprintf("%d@example.com", i+1)
	}
	fmt.Println("done sending  ")

	close(emailchan)

	<-done*/
	//emailchan <- "1@example.com"
	//emailchan <- "2@example.com"

	//fmt.Println(<-emailchan)
	//fmt.Println(<-emailchan)

	/*done := make(chan bool)
	go task(done)

	<-done*/

	/*result := make(chan int)

	go sum(result, 5, 10)
	res := <-result

	fmt.Println("result is", res)*/

	//numchan := make(chan int)

	//go processnum(numchan)

	//for {
	//	numchan <- rand.Intn(100)
	//}

	//numchan <- 5

	//time.Sleep(time.Second * 2)

	//messageChan := make(chan string)

	//messageChan <- "ping" //blocking

	//msg := <-messageChan

	//fmt.Println(msg)

}
