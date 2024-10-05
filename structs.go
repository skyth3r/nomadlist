package nomadlist

type Profile struct {
	Cached   string   `json:"cached"`
	Success  string   `json:"success"`
	Legal    string   `json:"legal"`
	Photo    string   `json:"photo"`
	Username string   `json:"username"`
	Location Location `json:"location"`
	Stats    Stats    `json:"stats"`
	Trips    []Trip   `json:"trips"`
}

type Location struct {
	Now      []any `json:"now"`
	Next     []any `json:"next"`
	Previous []any `json:"previous"`
}

type Stats struct {
	Cities                     int     `json:"cities"`
	Countries                  int     `json:"countries"`
	Followers                  int     `json:"followers"`
	Following                  int     `json:"following"`
	DistanceTraveledKM         int     `json:"distance_traveled_km"`
	DistanceTraveledMiles      int     `json:"distance_traveled_miles"`
	CountriesVisitedPercentage float64 `json:"countries_visited_percentage"`
	CitiesVisitedPercentage    float64 `json:"cities_visited_percentage"`
}

type Trip struct {
	EpochStart    int     `json:"epoch_start"`
	EpochEnd      int     `json:"epoch_end"`
	DateStart     string  `json:"date_start"`
	DateEnd       string  `json:"date_end"`
	Length        string  `json:"length"`
	EpochDuration int     `json:"epoch_duration"`
	Place         string  `json:"place"`
	PlaceSlug     string  `json:"place_slug"`
	PlaceLongSlug string  `json:"place_long_slug"`
	PlaceURL      string  `json:"place_url"`
	PlacePhoto    string  `json:"place_photo"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"country_code"`
	CountrySlug   string  `json:"country_slug"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	TripID        string  `json:"trip_id"`
}
