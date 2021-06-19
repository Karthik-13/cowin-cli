package api

type District struct {
	ID   int    `json:"district_id"`
	Name string `json:"district_name"`
}

type AllDistricts struct {
	Districts []District `json:"districts"`
	Ttl       int        `json:"ttl"`
}

func GetDistricts(stateId string) (a AllDistricts, err error) {

	req := createRequest("GET", PublicURLV2+"/admin/location/districts/"+stateId, nil, nil)
	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")
	req.Header.Set("Accept-Language", "hi_IN")
	err = sendRequest(req, &a)

	return a, err
}
