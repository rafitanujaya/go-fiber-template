package di

import "fmt"

func HealthCheck() {
	health := Injector.HealthCheck()
	fmt.Printf("Di HealthCheck: %v\n", health)
	isHealthy := true
	for service, err := range health {
		if err != nil {
			fmt.Printf("Service %s is unhealthy: %v\n", service, err)
			isHealthy = false
		}
	}
	if !isHealthy {
		panic("Di is not healthy")
	}
}
