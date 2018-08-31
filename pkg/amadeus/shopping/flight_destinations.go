package shopping

import (
	"github.com/alnacle/amadeus-golang/pkg/amadeus/client"
)

// FlightDestination represents the Flight Inspiration Search
type FlightDestination struct {
	Client *client.Client
}

func NewFlightDestination(client *client.Client) *FlightDestination {
	return &FlightDestination{client}
}

// Get JSON response
func (f *FlightDestination) Get(params ...string) ([]byte, error) {
	return f.Client.Get("/v1/shopping/flight-destinations", params...)
}
