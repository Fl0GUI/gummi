package gummi

import (
	"bufio"
	"fmt"
	"time"
)

var userInput *bufio.Scanner

func slowPrint(t string, perChar time.Duration) {
	for _, c := range t {
		fmt.Print(string(c))
		if !debug {
			time.Sleep(perChar)
		}
	}
}

func slowPrintLn(t string, perChar time.Duration) {
	slowPrint(t, perChar)
	fmt.Println()
}

func gummiSayTalk(t string, perChar time.Duration, after time.Duration) {
	fmt.Print("Gummi) ")
	slowPrintLn(t, perChar)
	if !debug {
		time.Sleep(after)
	}
}

func gummiSay(t string) {
	gummiSayTalk(t, time.Millisecond*25, time.Millisecond*200)
}

func gummiSaySlow(t string) {
	gummiSayTalk(t, time.Millisecond*500, time.Second)
}
