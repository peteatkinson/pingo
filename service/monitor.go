package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/rwxpeter/statusify/core"
)

type HeartbeatMonitor struct {
	config  *core.ServiceConfig
	handler func(core.ServiceConfig) StatusReport
}

func (m *HeartbeatMonitor) Tick(channel chan *HeartbeatMonitor) {
	var ticker *time.Ticker = time.NewTicker(time.Duration(m.config.HttpCallsFrequency) * time.Second)
	die := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			channel <- m
		case <-die:
			ticker.Stop()
			return
		}
	}
}

func (m *HeartbeatMonitor) CheckHeartbeat() {
	report := m.handler(*m.config)
	fmt.Println("Report", report)
	fmt.Println("Check heartbeat of ...", m.config.Url)
}

type StatusReport struct {
}

func GetHttpServiceHandler() func(core.ServiceConfig) StatusReport {
	return func(config core.ServiceConfig) StatusReport {
		fmt.Println("Running Http Service handler")
		return StatusReport{}
	}
}

func GetTcpServiceHandler() func(core.ServiceConfig) StatusReport {
	return func(config core.ServiceConfig) StatusReport {
		return StatusReport{}
	}
}

func GetStartTlsServiceHandler() func(core.ServiceConfig) StatusReport {
	return func(config core.ServiceConfig) StatusReport {
		return StatusReport{}
	}
}

func GetIcmpServiceHandler() func(core.ServiceConfig) StatusReport {
	return func(config core.ServiceConfig) StatusReport {
		return StatusReport{}
	}
}

func NewMonitor(config core.ServiceConfig) *HeartbeatMonitor {
	isServiceTCP := strings.HasPrefix(config.Url, "tcp://")
	isServiceICMP := strings.HasPrefix(config.Url, "icmp://")
	isServiceStartTLS := strings.HasPrefix(config.Url, "starttls://")
	isServiceHttp := !isServiceTCP && !isServiceICMP && !isServiceStartTLS

	if isServiceHttp {
		return &HeartbeatMonitor{config: &config, handler: GetHttpServiceHandler()}
	} else if isServiceICMP {
		return &HeartbeatMonitor{config: &config, handler: GetIcmpServiceHandler()}
	} else if isServiceStartTLS {
		return &HeartbeatMonitor{config: &config, handler: GetStartTlsServiceHandler()}
	} else {
		return nil
	}
}
