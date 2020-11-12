package util

import (
	consulapi "github.com/hashicorp/consul/api"
	"log"
)

func RegService() {
	config := consulapi.DefaultConfig()
	config.Address = "172.20.10.2:8500"

	reg := consulapi.AgentServiceRegistration{}
	reg.ID = "userservice"
	reg.Name = "userservice"
	reg.Address = "172.20.10.2"
	reg.Port = 8080
	reg.Tags = []string{"primary"}

	check := consulapi.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = "http://172.20.10.2:8080/health"

	reg.Check = &check

	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}
}
