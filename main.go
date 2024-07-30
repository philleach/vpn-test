package main

import (
	"fmt"
	"time"
)

var vpn Vpn_state = Vpn_state{connected: true}

func main() {
	vpn.start_pings("golang.org:http")
	for {
		fmt.Printf("Status :%v\n", vpn.get_state())
		time.Sleep(1 * time.Second)
	}
}
