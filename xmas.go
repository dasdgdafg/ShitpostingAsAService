package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

var linksInstance []string // static array of all the possible links
var once sync.Once

func getLinks() []string {
	once.Do(func() {
		f, e := os.Open("christmas.txt")
		if e != nil {
			log.Fatal(e)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			linksInstance = append(linksInstance, scanner.Text())
		}

		if e := scanner.Err(); e != nil {
			log.Fatal(e)
		}
		fmt.Println("christmas: found", len(linksInstance), "links")
	})
	return linksInstance
}

func xmas() string {
	links := getLinks()
	i := rand.Intn(len(links))
	timeLeft := getTimeUntil()
	return timeLeft + " " + links[i]
}

func getTimeUntil() string {
	now := time.Now()
	// we only care about the day
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	// this year's Christmas
	xmas := time.Date(now.Year(), 12, 25, 0, 0, 0, 0, time.Local)
	if now.Equal(xmas) {
		return "It's Christmas!"
	}
	if now.After(xmas) {
		// next year's Christmas
		xmas = time.Date(now.Year()+1, 12, 25, 0, 0, 0, 0, time.Local)
	}
	left := xmas.Sub(now)
	left = left.Round(24 * time.Hour)
	days := left.Hours() / 24
	dayString := "days"
	if days == 1 {
		dayString = "day"
	}
	return fmt.Sprintf("%v %s until Christmas", days, dayString)
}
