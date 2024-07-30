package main

import (
	"fmt"
	"time"
)

func main() {
	c := Vpn_state{connected: true}
	c.start_pings("golang.org:http")
	for {
		fmt.Printf("Status :%v\n", c.get_state())
		time.Sleep(1 * time.Second)
	}
}
