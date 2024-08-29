package resource

type CalLocationResource struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	NorthLat  float64 `json:"northLat"`
	SouthLat  float64 `json:"southLat"`
	EastLon   float64 `json:"eastLon"`
	WestLon   float64 `json:"westLon"`
}
