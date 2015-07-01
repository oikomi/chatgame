
package server

import (
	"sync"
	"time"
)

type HeartBeat struct {
	mu         sync.Mutex
	timeout    time.Duration
	expire     time.Duration
	fails      uint64
	threshold  uint64
}

func NewHeartBeat(timeout time.Duration, expire time.Duration, limit uint64) *HeartBeat {
	return &HeartBeat {
		timeout   : timeout,
		expire    : expire,
		threshold : limit,
	}
}

func (self *HeartBeat) ResetFailures() {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.fails = 0
}

func (self *HeartBeat) ChangeThreshold(thres uint64) {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.threshold = thres
}

func (self *HeartBeat) Beat() {
	timer := time.NewTicker(self.timeout * time.Second)
	//ttl := time.After(self.expire * time.Second)
	for {
		select {
		case <-timer.C:
			go func() {
				
			}()
		//case <-ttl:
			//break
		}
	}
}

func (self *HeartBeat) Receive() {
	timeout := time.After(self.timeout)
	for {
		select {
		case <-timeout:
			self.fails = self.fails + 1
			if self.fails > self.threshold {
				return
			}
		}
	}
}

