package amadeus

import (
	"github.com/alnacle/amadeus-golang/pkg/amadeus/client"
	"github.com/alnacle/amadeus-golang/pkg/amadeus/shopping"
)

// ShoppingStruct represents shopping namespace
type ShoppingStruct struct {
	Client            *client.Client
	FlightDestination *shopping.FlightDestination
	FlightOffers      *shopping.FlightOffers
}

func NewShopping(client *client.Client) *ShoppingStruct {
	return &ShoppingStruct{
		Client:            client,
		FlightDestination: &shopping.FlightDestination{client},
		FlightOffers:      &shopping.FlightOffers{client},
	}
}
