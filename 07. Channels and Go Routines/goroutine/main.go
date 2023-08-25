package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://facebook.com",
		"https://google.com",
		"https://golang.org",
		"https://stackoverflow.com",
	}

	c := make(chan string) //channel creation
	for _, link := range links {
		//by default every program in Go has a "1 main routine" that compiles the code
		go checkURL(link, c) //"go" keyword will assign a new "goroutine"(aka child routine) and the "go scheduler" will manage/handle the Goroutines(threads) internally
		//Channels is the only way to communicate between different goroutines
		//without a Channel,(lets take our code example), the main routine(which is responsible for quiting our program) will close after it assigns the tasks to its child goroutine, it doesn't care if the child goroutines have finished executing or not, as a result unexpected results are seen, we do need channels here for the communication
		// (in our code example)after the loop has exited and all the goroutines have been created, we are expectecting fmt.Println(<-c) (only single statement),(because this staetment is a blocking code) so as soon as a message is recieved from any of the goroutines, the message is sent to this channel and the proggram terminates(not what we want, we want all goroutines to finish its message to sending )
	}
	//fmt.Println(<-c) //recieving a channel message is a blocking operation, if we only recieve 1 message, our main routine will terminate after printing 1 message

	//fmt.Println(<-c) //for that reason, one way (a very bad way) is to recieve exactly the same number of message as we have sent(in our case 4 messages for 4 links)
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	//or obvious way is to use the loop
	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c)
	//}

	//checking for links status, continuosly forever
	//for {
	//	go checkURL(<-c, c)
	//}

	//OR
	//for l := range c { //wait for the channel to recieve the value, if it  has recieved, then store it to "l"
	//	go checkURL(l, c)
	//}

	//lets make the link calls every 5 seconds, rather than continuously checking the links very fast
	for l := range c {
		//time.Sleep(5 * time.Second) //not a proper place to put the put the time.Sleep() method, as it will just pause the main routine and underway all the child goroutines which might have finished execution are waiting to send message to the channel (a very bad design, as it just stacks up the goroutines messages which have finished its execution and waiting to be recieved by the channel), instead use the time.Sleep() for the child goroutines
		//go checkURL(l, c)

		//another way using function literal
		go func(link string) { //NOTE: this function literal is same as Anonymous function in javascript and we are immediately invoking the function, which is same as IIFE(immediaterly invoking function expression) in js
			time.Sleep(5 * time.Second)
			checkURL(link, c)
		}(l) //using the copy value of the outer func "l", to solve the same varible pointing issue in goroutine (never share the same variable of main routine into its child goroutine)
	}
}

func checkURL(link string, c chan string) {
	//time.Sleep(5*time.Second) //we can sure keep the time.Sleep() method here, it will not stack up the goroutines message waiting to be put on the channel, BUT there is one more relevant place, using the function literal
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!!")
		//c <- "Link might be down(from channel)"
		c <- link
		return
	}
	fmt.Println(link, "is up!!")
	//c <- "Link is up(from channel)"
	c <- link
}
