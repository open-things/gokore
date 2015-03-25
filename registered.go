package gokore

import (
	"fmt"
	"sort"
)

var registered Endpoints = nil

func RegisterEndpoint(endpoint EndpointInterface) {
	registered = append(registered, endpoint)
	sort.Sort(EndpointsByTitle(registered))
}

func GetEndpointCount() int { return len(registered) }
func GetEndpointByIndex(index int) (EndpointInterface, error) {
	if 0 <= index && index < len(registered) {
		return registered[index], nil
	} else {
		return nil, fmt.Errorf("Index is out of bounds.")
	}
}
func GetEndpointByTitle(title string) (EndpointInterface, error) {
	for _, endpoint := range registered {
		if endpoint.Title() == title {
			return endpoint, nil
		}
	}
	return nil, fmt.Errorf("Title not found")
}
