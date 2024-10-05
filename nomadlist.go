package nomadlist

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Profile(username string) (*Profile, error) {
	url := fmt.Sprintf("https://nomadlist.com/@%s.json", username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var profile Profile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}

	return &profile, nil
}

func (c *Client) Stats(username string) (*Stats, error) {
	profile, err := c.Profile(username)
	if err != nil {
		return nil, err
	}

	return &profile.Stats, nil
}

func (c *Client) Location(username string) (*Location, error) {
	profile, err := c.Profile(username)
	if err != nil {
		return nil, err
	}

	return &profile.Location, nil
}

func (c *Client) Trips(username string) ([]Trip, error) {
	profile, err := c.Profile(username)
	if err != nil {
		return nil, err
	}

	return profile.Trips, nil
}

func (c *Client) Trip(username, tripID string) (*Trip, error) {
	trips, err := c.Trips(username)
	if err != nil {
		return nil, err
	}

	for _, trip := range trips {
		if trip.TripID == tripID {
			return &trip, nil
		}
	}

	return nil, fmt.Errorf("trip with id %s not found", tripID)
}

func (c *Client) TripsInYear(username, year string) ([]Trip, error) {
	trips, err := c.Trips(username)
	if err != nil {
		return nil, err
	}

	var tripsInYear []Trip
	for _, trip := range trips {
		if trip.DateStart[:4] == year {
			tripsInYear = append(tripsInYear, trip)
		}
	}

	return tripsInYear, nil
}

func (c *Client) TripsInCountry(username, country string) ([]Trip, error) {
	trips, err := c.Trips(username)
	if err != nil {
		return nil, err
	}

	var tripsInCountry []Trip
	for _, trip := range trips {
		if trip.Country == country {
			tripsInCountry = append(tripsInCountry, trip)
		}
	}

	return tripsInCountry, nil
}

func (c *Client) TripsInCity(username, city string) ([]Trip, error) {
	trips, err := c.Trips(username)
	if err != nil {
		return nil, err
	}

	var tripsInCity []Trip
	for _, trip := range trips {
		if trip.Place == city {
			tripsInCity = append(tripsInCity, trip)
		}
	}

	return tripsInCity, nil
}
