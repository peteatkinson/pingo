package main

import (
	"fmt"

	"github.com/rwxpeter/statusify/core"
	"github.com/rwxpeter/statusify/service"
)

var (
	channel chan *service.HeartbeatMonitor
)

func main() {
	// schedule("Test 1", 60)
	// schedule("Test 2", 60)

	serviceConfigs := core.LoadServiceConfig()

	if len(serviceConfigs) > 0 {
		channel = make(chan *service.HeartbeatMonitor, len(serviceConfigs))
		TickOverHeartbeatMonitors(serviceConfigs)
	}

	fmt.Scanln()
}

func TickOverHeartbeatMonitors(schedules []core.ServiceConfig) {
	go SetupHeartbeatListeners()

	for _, config := range schedules {
		//schedule(config)
		monitor := service.NewMonitor(config)
		go monitor.Tick(channel)
	}
}

// func handleSchedule(config core.ServiceConfig) {
// 	fmt.Println(config)
// 	fmt.Println(service.GetServiceStatus(config.Url, time.Duration(config.DurationThreshold) * time.Millisecond))
// }

func SetupHeartbeatListeners() {
	for {
		select {
		case monitor := <-channel:
			go monitor.CheckHeartbeat()
		}
	}
}
