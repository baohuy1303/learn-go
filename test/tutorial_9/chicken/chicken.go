package main

import (
	"fmt"
	"math/rand"
	"time"

)

var maxChickenPrice float32 = 5
var maxTofuPrice float32 = 3

// FIND THE FIRST DEAL FROM ALL 3 SOURCES

func main() {
	var timeStart = time.Now()
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart", "target", "costco"}
	var doneChicken = make(chan bool)
	var doneTofu = make(chan bool)
	for i := range websites {
		go checkWebsiteForChicken(websites[i], chickenChannel, doneChicken)
		go checkWebsiteForTofu(websites[i], tofuChannel, doneTofu)
	}
	sendMessage(chickenChannel, tofuChannel, doneChicken, doneTofu, timeStart)
}

func checkWebsiteForChicken(website string, chickenChannel chan string, done chan bool) {
	for{
		select{
			case <- done:
				return
			default:
				time.Sleep(1 * time.Second)
				var chickenPrice = rand.Float32()*20
				if chickenPrice <= maxChickenPrice{
					chickenChannel <- website
					// no sending done to true here because none is reading and will
					// create a deadlock
					return
				}
		}
	}
}

func checkWebsiteForTofu(website string, tofuChan chan string, done chan bool) {
	for{
		select{
			case <- done:
				return
			default:
				time.Sleep(1 * time.Second)
				var tofuPrice = rand.Float32()*20
				if tofuPrice <= maxTofuPrice{
					tofuChan <- website
					// no sending done to true here because none is reading and will
					// create a deadlock
					return
				}
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChan chan string, doneChicken chan bool,
	doneTofu chan bool, timeStart time.Time){
	for{
	select{
	case m:= <- chickenChannel:
			fmt.Printf("Found a chicken deal at %s\n", m)
			close(doneChicken)
	case m:= <- tofuChan:
			fmt.Printf("Found a tofu deal at %s\n", m)
			close(doneTofu)
		}
	fmt.Println(time.Since(timeStart))
	}

	//SEE HOW IT CANT BE CLOSED properly. That's where struct comes in
	// to handle all these cleanly.
}