package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/micro/mdns"
)

func getPairingCode(port int) (string, error) {
	ips, err := getLocalIPs()
	if err != nil {
		return "", err
	}

	ipStrings := []string{}
	for _, ip := range ips {
		ipStrings = append(ipStrings, ip.String())
	}

	return fmt.Sprintf("%d,%s", port, strings.Join(ipStrings, ",")), nil
}

func getLocalIPs() ([]net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	ips := []net.IP{}
	for _, iface := range ifaces {
		// Do not include if interface:
		//   - Is not up
		//   - Is a loopback address
		if iface.Flags&net.FlagUp != net.FlagUp ||
			iface.Flags&net.FlagLoopback == net.FlagLoopback {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			switch ip := addr.(type) {
			case *net.IPNet:
				if ip.IP.DefaultMask() == nil {
					continue
				}

				ips = append(ips, ip.IP)
			}
		}
	}

	return ips, nil
}

func advertiseService(port int) *mdns.Server {
	hostname, _ := os.Hostname()
	pc, _ := getPairingCode(port)
	info := []string{"version=1", fmt.Sprintf("pairing=%s", pc)}
	service, _ := mdns.NewMDNSService(hostname, "_vrhr._tcp", "", "", port, nil, info)

	server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	return server
}
