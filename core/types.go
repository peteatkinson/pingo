package core

import (
	"fmt"
	"time"
)

type ServiceStatus struct {
	Duration   time.Duration
	StatusCode int
	HealthStatus string
}

type ServiceConfig struct  {
	Enabled 						*bool `json:"enabled"`
	Url       					string `json:"url"`
	HttpCallsFrequency 	int `json:"frequency"`
	DurationThreshold 	int `json:"duration_threshold"`
}

const (
	ServiceDead = "SERVICE_DEAD"
	ServiceHealthy = "SERVICE_HEALTHY"
	ServiceUnhealthy = "SERVICE_UNHEALTHY"
)

func NewUnhealthyServiceStatus(duration time.Duration, statusCode int) (ServiceStatus) {
	return ServiceStatus{Duration: duration, StatusCode: statusCode, HealthStatus: ServiceUnhealthy}
}

func NewHealthyServiceStatus(duration time.Duration, statusCode int) (ServiceStatus) {
	return ServiceStatus{Duration: duration, StatusCode: statusCode, HealthStatus: ServiceHealthy}
}

func NewServiceDeadServiceStatus() (ServiceStatus) {
	return ServiceStatus{Duration: 0, StatusCode: 0, HealthStatus: ServiceDead}
}

func CheckHealthStatus(duration time.Duration, durationThreshold time.Duration, statusCode int) (string) {
	if statusCode == 0 {
		return ServiceDead
	}
	fmt.Println("threshold: ", durationThreshold, duration)
	if duration > durationThreshold {
		return ServiceUnhealthy
	}

	return ServiceHealthy
}