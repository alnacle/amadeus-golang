package main

import (
	"fmt"

	"github.com/alnacle/amadeus-golang/pkg/amadeus"
)

func main() {

	client := amadeus.Create()
	fmt.Println(client.Version)
	fmt.Printf("%+v\n", client)

	response, _ := client.Shopping.FlightDestination.Get("origin=MAD", "maxPrice=200")
	fmt.Println(string(response))

}
