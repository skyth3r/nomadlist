# Nomadlist
A collection of useful functions to get your own travel stats from your [Nomadlist](https://nomadlist.com/) profile

**Status:** Work in Progress

## Getting Started

### Installing
Use `go get` to retrieve the package to add it to your project's Go module dependencies.

	go get github.com/skyth3r/nomadlist

To update the package use `go get -u` to retrieve the latest version of the package.

	go get -u github.com/skyth3r/nomadlist

## Docs

🔜

## Quick Examples

```Go
// Create new client
client := nomadlist.NewClient()

// Get user's profile
profile, err := client.profile("USERNAME")
if err != nil {
    fmt.Println(err)
}
fmt.Println(diary)

// Get all trips
trips, err := client.Trips("USERNAME")
if err != nil {
	fmt.Println(err)
	os.Exit(1)
}
fmt.Println(trips)

// Get trips in a certain year
tripsIn2023, err := client.TripsInYear("USERNAME", "2023")
if err != nil {
	fmt.Println(err)
	os.Exit(1)
}
fmt.Println(tripsIn2023)
```