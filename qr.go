package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/mdp/qrterminal/v3"
)

func printQrInfo() {
	fmt.Printf("You can pair your phone to this server by scanning this QR code:\n\n")

	ips, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Nevermind. Something went wrong while trying to find your machine's IP address.")
		return
	}

	// Trust that the first interface is local network
	ip := strings.Split(ips[0].String(), "/")[0]
	qrterminal.Generate(fmt.Sprintf("http://%s%s", ip, envListenAddr), qrterminal.L, os.Stdout)

	fmt.Printf("\nIf the QR code doesn't work, make sure your phone and this machine are on the same network.\nIf it still doesn't work after that, check your machine's settings to see if %s is the IP for your local network.\n", ip)
}
