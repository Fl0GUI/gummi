package gummi

import (
	"fmt"
	"log"
	"time"

	"j322.ica/gumroad-sammi/validate"
)

var debug = false

var userName string

func printHello() {
	log.Println("First time setup")
	if !debug {
		time.Sleep(time.Second)
	}
	fmt.Print("Gummi assistant: starting")
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

func printFailure() {
	if !debug {
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Print(gummiSquint)
	if !debug {
		time.Sleep(time.Millisecond * 200)
	}
	gummiSay("There still seems to be a problem, let's keep going.")
	if !debug {
		time.Sleep(time.Millisecond * 200)
	}
}

func printIssue(f *validate.Functions) {
	if f.Valid() {
		return
	}

	fmt.Print(gummiSquint)
	gummiSay("I did my best to set things up automatically, but I encountered some issues.")
	if f.Sammi != nil {
		gummiSay("I could not connect to sammi.")
	}
	if f.Gumroad != nil {
		gummiSay("I could not set up my gumroad connection.")
	}
	if f.Throne != nil {
		gummiSay("I could not connect to throne's database.")
	}

	if f.Sammi != nil && f.Gumroad != nil && f.Throne != nil {
		gummiSay("Lets look at the sammi connection first.")
	}
}
