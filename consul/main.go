package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

// RegisterService registers a service with Consul
func RegisterService(serviceID, serviceName, serviceAddress string, servicePort int, healthCheckPath string) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatalf("Error creating Consul client: %v", err)
	}

	// Construct the full health check URL
	healthCheckURL := fmt.Sprintf("http://%s:%d%s", serviceAddress, servicePort, healthCheckPath)

	serviceRegistration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Address: serviceAddress,
		Port:    servicePort,
		Tags:    []string{"gRPC", "HTTP"},
		Check: &api.AgentServiceCheck{
			HTTP:     healthCheckURL, // Fully qualified health check URL
			Interval: "10s",          // Health check interval
			Timeout:  "5s",           // Health check timeout
		},
	}

	// Register the service with Consul
	err = client.Agent().ServiceRegister(serviceRegistration)
	if err != nil {
		log.Fatalf("Failed to register %s: %v", serviceName, err)
	}

	log.Printf("%s registered successfully with Consul!", serviceName)
}

func main() {
	// Register BlogService
	RegisterService(
		"blogservice-1", // Unique service ID
		"blogservice",   // Service name
		"127.0.0.1",     // Service address
		8081,            // Service port
		"/health",       // Health check path
	)

	// Register UserService
	RegisterService(
		"userservice-1",
		"userservice",
		"127.0.0.1",
		8082,
		"/health",
	)

	// Register AuthService
	RegisterService(
		"authservice-1",
		"authservice",
		"127.0.0.1",
		8083,
		"/health",
	)
}
