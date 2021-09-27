package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

var (
	bus chan int
)

func main() {
	// schedule("Test 1", 60)
	// schedule("Test 2", 60)
	loadSchedule()
	fmt.Scanln()
}

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

type Schedule struct {
	Url string `json:"url"`
}

func loadSchedule() {
	var schedule []Schedule

	data, err := ioutil.ReadFile("./schedule.json")

	if err != nil {
		fmt.Println("Error reading JSON file")
	}

	err = json.Unmarshal(data, &schedule)
	if err != nil {
		fmt.Println("Error unmarshing JSON file")
	}
	fmt.Println(schedule)
}

func schedule(name string, seconds int) {
	var ticker *time.Ticker = time.NewTicker(5 * time.Second)
	die := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println(name, ticker.C)
			case <-die:
				ticker.Stop()
				return
			}
		}
	}()
}
