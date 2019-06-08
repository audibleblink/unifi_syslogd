# UniFi Syslogd

Simple syslog server to receive firewall logs from USG/Edge Routers

## Usage

1. Enable logging on the firewall rules you want to troubleshoot

![unifi firewall config](https://i.imgur.com/x2wlxy5.png)

2. Enable offloading of logs to a syslog server under Settings > Site

![unifi syslog config](https://i.imgur.com/IJlSE1S.png)

```
# build with
go build -o unifi_syslog main.go
./unifi_syslog
```
