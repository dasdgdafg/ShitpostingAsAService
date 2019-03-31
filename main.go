package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/dasdgdafg/ircFramework"
	"github.com/dasdgdafg/markovChain"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const botName = "ShitpostingAsAService"
const ident = "always"
const realname = "dasdgdafg"

var chains = make(map[string]*markovChain.Chain)
var shitpostRegex = regexp.MustCompile("^(\\d,?\\d?)?(>)?!shitpost$")
var s_h_i_t_p_o_s_t_R_e_g_e_x = regexp.MustCompile("^(\\d,?\\d?)?(>)? ?! ?s h i t p o s t ?$")

var messages = make(map[string][]string)
var sedRegex = regexp.MustCompile("^s/.*/.*/.*$")

var xmasRegex = regexp.MustCompile("^(\\d,?\\d?)?(>)?!(?:c|C)hristmas$")

var diceRegex = regexp.MustCompile("^!(\\d*)d(\\d+)$")

func processLine(linesToSend chan<- string, nick string, channel string, msg string) {
	if chains[channel] != nil {
		c := chains[channel]
		var matches []string
		long := false
		if shitpostRegex.MatchString(msg) {
			matches = shitpostRegex.FindStringSubmatch(msg)
		} else if s_h_i_t_p_o_s_t_R_e_g_e_x.MatchString(msg) {
			matches = s_h_i_t_p_o_s_t_R_e_g_e_x.FindStringSubmatch(msg)
			long = true
		}
		if matches != nil {
			line := c.Generate(100)
			if long {
				line = insertNth(line, 1)
			}
			if len(matches) > 1 {
				toAdd := matches[1:]
				line = strings.Join(toAdd, "") + line
			}
			linesToSend <- "PRIVMSG " + channel + " :" + line
		}
	}
	if strings.HasPrefix(channel, "#") {
		if sedRegex.MatchString(msg) {
			// run the command in a separate go routine
			go sed(messages[channel], msg, channel, linesToSend)
		}
		messages[channel] = append(messages[channel], "<"+nick+"> "+msg)
		if len(messages[channel]) > 10 {
			messages[channel] = messages[channel][1:]
		}
	}
	if xmasRegex.MatchString(msg) {
		matches := xmasRegex.FindStringSubmatch(msg)
		line := xmas()
		if len(matches) > 1 {
			toAdd := matches[1:]
			line = strings.Join(toAdd, "") + line
		}
		linesToSend <- "PRIVMSG " + channel + " :" + line
	}
	if diceRegex.MatchString(msg) {
		line := diceResult(msg)
		linesToSend <- "PRIVMSG " + channel + " :" + line
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	prefixLen := 5

	// generate markov chains for each channel that we have input for
	files, err := ioutil.ReadDir("input")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println("generating markov chain for " + file.Name())

		c := markovChain.NewChain(prefixLen)

		inputFile, err := os.Open("input/" + file.Name())
		if err != nil {
			panic(err)
		}

		avoidFile, err := os.Open("avoid/" + file.Name())
		if err != nil {
			panic(err)
		}

		c.Build(bufio.NewReader(inputFile), bufio.NewReader(avoidFile))
		inputFile.Close()
		avoidFile.Close()
		chains[file.Name()] = c
	}

	passwordBytes, err := ioutil.ReadFile("password.txt")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	password := string(passwordBytes)

	bot := ircFramework.IRCBot{Server: "irc.rizon.net",
		Port:           "6697",
		Nickname:       botName,
		Ident:          ident,
		Realname:       realname,
		Password:       password,
		Statefile:      "shitpost.state",
		ListenToStdin:  false,
		MessageHandler: processLine,
	}
	bot.Run()
}

func diceResult(msg string) string {
	matches := diceRegex.FindStringSubmatch(msg)
	num, err := strconv.ParseInt(matches[1], 10, 64)
	if matches[1] == "" {
		num = 1
	} else if err != nil {
		return "Unable to parse int64 from " + matches[1]
	}
	sides, err := strconv.ParseInt(matches[2], 10, 64)
	if err != nil {
		return "Unable to parse int64 from " + matches[2]
	}
	if sides == 0 {
		return "0"
	}
	return rollDice(sides, num)
}

//https://stackoverflow.com/questions/33633168/how-to-insert-a-character-every-x-characters-in-a-string-in-golang
func insertNth(s string, n int) string {
	var buffer bytes.Buffer
	var n_1 = n - 1
	var l_1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n_1 && i != l_1 {
			buffer.WriteRune(' ')
		}
	}
	return buffer.String()
}
