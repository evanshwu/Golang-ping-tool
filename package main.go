package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
	// Network library for ping tool
	"github.com/sparrc/go-ping"
)

func main() {
	// Set input scanner
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the ip to ping: ")
	scanner.Scan()
	inputIp := scanner.Text()

	// init ping utility
	pinger, err := ping.NewPinger(strings.TrimSpace(inputIp))
	pinger.SetPrivileged(true)
	pinger.OnFinish = func(stats *ping.Statistics) {
		// print packet loss and latency
		fmt.Printf("%v%% packet loss, %v latency\n", math.Round(stats.PacketLoss), stats.Rtts)
	}

	// Error handling
	if err != nil {
		panic(err)
	}

	// emulates a while loop to ping every duration milliseconds
	var input int
	for ok := true; ok; ok = (input != 2) {
		pinger.Count = 1
		pinger.Run()                 // blocks until finished
		stats := pinger.Statistics() // get send/receive/rtt stats
		fmt.Println(stats)
		// Sleeps for one second to slow down the ping frequncy
		time.Sleep(1000 * time.Millisecond)
	}

}
