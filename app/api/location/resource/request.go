package resource

type FindEventRequest struct {
	Latitude  string `json:"latitude" query:"latitude"`
	Longitude string `json:"longitude" query:"longitude"`
}
