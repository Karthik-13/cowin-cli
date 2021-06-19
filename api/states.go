package api

type State struct {
	ID   int    `json:"state_id"`
	Name string `json:"state_name"`
}

type AllStates struct {
	States []State `json:"states"`
	Ttl    int64   `json:"ttl"`
}

func GetStates() (a AllStates, err error) {

	req := createRequest("GET", "cdn-api.co-vin.in/api/v2/admin/location/states", nil, nil)
	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

	err = sendRequest(req, &a)

	return a, err
}
