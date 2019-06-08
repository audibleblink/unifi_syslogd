package main

import "fmt"
import "gopkg.in/mcuadros/go-syslog.v2"

func main() {
	server := syslog.NewServer()
	channel := make(syslog.LogPartsChannel)

	server.SetFormat(syslog.RFC3164)
	server.SetHandler(syslog.NewChannelHandler(channel))

	server.ListenUDP("0.0.0.0:514")
	server.Boot()

	go func(ch syslog.LogPartsChannel) {
		for logParts := range ch {
			fmt.Println(logParts)
		}
	}(channel)

	server.Wait()
}
