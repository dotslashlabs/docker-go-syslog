package main

import (
	"fmt"
	"os"

	"gopkg.in/mcuadros/go-syslog.v2"
)

func main() {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.RFC3164)
	server.SetHandler(handler)
	server.SetTimeout(10)

	addrUDP := fmt.Sprintf("%s:%s", getenv("HOST_UDP", getenv("HOST", "0.0.0.0")), getenv("PORT_UDP", "514"))
	err := server.ListenUDP(addrUDP)
	check(err)

	addrTCP := ""
	portTCP := getenv("PORT_TCP", "")

	if portTCP != "" {
		addrTCP = fmt.Sprintf("%s:%s", getenv("HOST_TCP", getenv("HOST", "0.0.0.0")), portTCP)
		err = server.ListenTCP(addrTCP)
		check(err)
	}

	err = server.Boot()
	check(err)

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			fmt.Println(logParts)
		}
	}(channel)

	s1 := fmt.Sprintf("Syslog server listening on udp:%s", addrUDP)
	if addrTCP != "" {
		s1 = fmt.Sprintf("%s tcp:%s", s1, addrTCP)
	}

	fmt.Println(s1, "...")
	server.Wait()
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
