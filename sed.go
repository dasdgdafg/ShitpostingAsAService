package main

import (
	"math/rand"
	"os/exec"
	"strings"
)

func sed(strs []string, script string, channel string, sedOutput chan<- string) {
	// memes
	if rand.Intn(10) == 0 {
		script = "s/>.*/> HAHAHA DISREGARD THAT, I SUCK COCKS/"
	}

	// I think this is secure...
	cmd := exec.Command("sed", "--sandbox", "-e", script)
	input, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	for _, str := range strs {
		input.Write([]byte(str + "\n"))
	}
	input.Close()
	out, err := cmd.CombinedOutput()

	// print errors to the channel
	if err != nil {
		lines := strings.Replace(string(out), "\n", " ", -1)
		sedOutput <- "PRIVMSG " + channel + " :" + lines
		return
	}

	// TODO: handle lines getting turned into multiple lines
	outStrs := strings.Split(string(out), "\n")
	oLen := len(outStrs)
	iLen := len(strs)
	if oLen > iLen {
		// chop off extra \n at the end
		outStrs = outStrs[:iLen]
		oLen = iLen
	}
	// print first thing that changed
	for i := 1; i <= oLen && i <= iLen; i++ {
		if outStrs[oLen-i] != strs[iLen-i] {
			sedOutput <- "PRIVMSG " + channel + " :" + outStrs[oLen-i]
			return
		}
	}
}
