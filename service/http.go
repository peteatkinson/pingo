package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rwxpeter/statusify/core"
)


func GetServiceStatus(url string, threshold time.Duration) (s core.ServiceStatus) {
	timeStart := time.Now()

	response, err := execute(url)

	if err != nil {
		fmt.Println(response)
		return core.NewServiceDeadServiceStatus();
	}

	duration := time.Since(timeStart)
	defer response.Body.Close()

	healthStatus := core.CheckHealthStatus(duration, threshold, response.StatusCode)

	if healthStatus == core.ServiceHealthy {
		return core.NewHealthyServiceStatus(duration, s.StatusCode)
	}

	if healthStatus == core.ServiceUnhealthy {
		return core.NewUnhealthyServiceStatus(duration, s.StatusCode)
	}

	return core.NewServiceDeadServiceStatus()
}

func execute(url string) (*http.Response, error){
	client := getHttpClient()
	response, err := client.Do(createRequest(url))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func getHttpClient() *http.Client {
	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	return client
}

func createRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	return req
}