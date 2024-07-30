package main

import (
	"net"
	"sync"
	"time"
)

type Vpn_state struct {
	mu        sync.Mutex
	connected bool
}

func (c *Vpn_state) set_state(state bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.connected = state
}

func (c *Vpn_state) get_state() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.connected
}

func (c *Vpn_state) ping_host(host string) {
	conn, err := net.DialTimeout("tcp", host, 1*time.Second)

	if err != nil {
		c.set_state(false)
	} else {
		defer conn.Close()
		c.set_state(true)
	}
}

func (c *Vpn_state) start_pings(host string) {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				c.ping_host(host)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
