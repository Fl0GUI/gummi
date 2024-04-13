package gummi

import (
	"fmt"
	"log"
	"time"
)

var debug = false

var userName string

func printHello() {
	log.Println("first time setup")
	if !debug {
		time.Sleep(time.Second)
	}
	fmt.Print("Starting Gummi assistant")
	if !debug {
		time.Sleep(time.Second)
	}
	slowPrintLn("........", time.Millisecond*200)

	fmt.Print(gummiSmile)
	if !debug {
		time.Sleep(time.Second)
	}
	gummiSay("Hello, I am Gummi! I am here to help you set up this application.")
	gummiSay("How are you today?")
	if !debug {
		prompt()
	}
}

func printSucceeded() {
	fmt.Print(gummiYay)
	gummiSay("YAY! Everything is looking good. Thank you for helping me!")
	gummiSay("I will save these configs for next time.")
	gummiSaySlow(".....")
	fmt.Print(gummiSad)
	gummiSay("That does mean I will not be around next time.")
	gummiSay("But if you ever want to work together again you can delete my configuration file.")
	fmt.Print(gummiSmile)
	gummiSay("With that said, I will be shutting down now. I wish you the best of luck and goodbye.")
	if !debug {
		time.Sleep(time.Millisecond * 1000)
	}
	fmt.Print(gummiClose)
	if !debug {
		time.Sleep(time.Millisecond * 1000)
	}
}

func printIssue(f *functions) {
	if f.valid() {
		return
	}

	fmt.Print(gummiSquint)
	gummiSay("I did my best to set things up automatically, but I encountered some issues.")
	if !f.validSammi() {
		gummiSay("I could not connect to your sammi button.")
	}
	if !f.validGumroad() {
		gummiSay("I could not set up my gumroad connection.")
	}
	if !f.validSammi() && !f.validGumroad() {
		gummiSay("Lets look at the sammi connection first.")
	}
}
