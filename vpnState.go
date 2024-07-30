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

func (v *Vpn_state) set_state(state bool) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.connected = state
}

func (v *Vpn_state) get_state() bool {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.connected
}

func (v *Vpn_state) ping_host(host string) {
	conn, err := net.DialTimeout("tcp", host, 1*time.Second)

	if err != nil {
		v.set_state(false)
	} else {
		defer conn.Close()
		v.set_state(true)
	}
}

func (v *Vpn_state) start_pings(host string) {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				v.ping_host(host)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
