package main

import (
	"math"
	"net"
	"os/exec"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/digineo/go-ping"
)

func main() {
	// Create new pinger
	pinger, err := ping.New("0.0.0.0", "")
	if err != nil {
		panic(err)
	}

	// Resolve domain or IP address
	ip, err := net.ResolveIPAddr("ip4", "8.8.8.8")
	if err != nil {
		panic(err)
	}

	// Execute pinger with retry
	retry.Retry(func(attempt uint) error {
		var err error
		for err == nil {
			_, err := pinger.Ping(ip, 5*time.Second)
			if err != nil {
				// Execute exec.bat on error
				cmd := exec.Command("cmd.exe", "/C", "exec.bat")
				cmd.Run()
				return err
			}
			// Wait 5 seconds until next ping
			time.Sleep(5 * time.Second)
		}
		return err
	}, strategy.Limit(math.MaxInt64), strategy.Wait(30*time.Second))
	if err != nil {
		panic(err)
	}
}
