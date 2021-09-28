package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	bus chan string
)

func main() {
	// schedule("Test 1", 60)
	// schedule("Test 2", 60)

	schedules, err := loadSchedulesConfig()

	if err != nil {
		os.Exit(0)
	}

	bus = make(chan string, len(schedules))
	runSchedules(schedules)

	fmt.Scanln()
}

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

type Schedule struct {
	Url       string `json:"url"`
	Frequency int    `json:"frequency"`
}

func runSchedules(schedules []Schedule) {
	go handleSchedulers()

	for _, s := range schedules {
		schedule(s.Url, s.Frequency)
	}
}

func loadSchedulesConfig() ([]Schedule, error) {
	var schedules []Schedule

	data, err := ioutil.ReadFile("./schedule.json")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &schedules)

	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func schedule(url string, frequency int) {
	var ticker *time.Ticker = time.NewTicker(time.Duration(frequency) * time.Second)
	die := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Tick!", url)
				bus <- url
			case <-die:
				ticker.Stop()
				return
			}
		}
	}()
}

func handleSchedule(url string) {
	fmt.Println(url)
}

func handleSchedulers() {
	for {
		select {
		case url := <-bus:
			go handleSchedule(url)
		}
	}
}
