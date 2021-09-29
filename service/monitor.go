package service

import (
	"fmt"
	"time"

	"github.com/rwxpeter/statusify/core"
)

type HeartbeatMonitor struct {
	config *core.ServiceConfig
}


func (m *HeartbeatMonitor) Tick(channel chan *HeartbeatMonitor) {
	var ticker *time.Ticker = time.NewTicker(time.Duration(m.config.HttpCallsFrequency) * time.Second)
	die := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			channel <-m
		case <-die:
			ticker.Stop()
			return
		}
	}
}

func (m *HeartbeatMonitor) CheckHeartbeat() {
	fmt.Println("Check heartbeat of ...", m.config.Url)
}


func NewMonitor(config core.ServiceConfig) (HeartbeatMonitor) {
	return HeartbeatMonitor{config: &config}
}