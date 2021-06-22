package api

import (
	"net/url"
)

type AllCenters struct {
	Centers []Centers `json:"centers"`
}

type CenterByCenter struct {
	Centers Centers `json:"centers"`
}

type Centers struct {
	CenterId       int    `json:"center_id"`
	CenterName     string `json:"name"`
	CenterState    string `json:"state_name"`
	CenterDistrict string `json:"district_name"`
	CenterBlock    string `json:"block_name"`
	CenterPincode  int    `json:"pincode"`
	FeeType        string `json:"fee_type"`
	//VaccineSessions       json.RawMessage `json:"sessions"`
	Sessions []Sessions `json:"sessions"`
}

type Sessions struct {
	Date             string   `json:"date"`
	AvailableCapcity int      `json:"available_capacity"`
	AgeLimit         int      `json:"min_age_limit"`
	VaccineName      string   `json:"vaccine"`
	FirstDose        int      `json:"available_capacity_dose1"`
	SecondDose       int      `json:"available_capacity_dose2"`
	AvailableSlots   []string `json:"slots"`
}

func GetCentersSessionsByPin(pincode string, appointmentDate string, vaccine string) (s AllCenters, err error) {
	q := url.Values{}
	q.Add("pincode", pincode)
	q.Add("date", appointmentDate)
	q.Add("vaccine", vaccine)
	req := createRequest("GET", PublicURLV2+"/appointment/sessions/public/calendarByPin", q, nil)
	req.Header.Set("User-Agent", "Cowin_Cli/1.0")
	req.Header.Set("Accept-Language", "hi_IN")
	err = sendRequest(req, &s)
	return s, err
}

func GetCentersSessionsByDistrict(districtId string, appointmentDate string, vaccine string) (s AllCenters, err error) {
	q := url.Values{}
	q.Add("district_id", districtId)
	q.Add("date", appointmentDate)
	q.Add("vaccine", vaccine)
	req := createRequest("GET", PublicURLV2+"/appointment/sessions/public/calendarByDistrict", q, nil)
	req.Header.Set("User-Agent", "Cowin_Cli/1.0")
	req.Header.Set("Accept-Language", "hi_IN")
	err = sendRequest(req, &s)
	return s, err
}

func GetCentersSessionsByCenter(centerId string, appointmentDate string, vaccine string) (s CenterByCenter, err error) {
	q := url.Values{}
	q.Add("center_id", centerId)
	q.Add("date", appointmentDate)
	q.Add("vaccine", vaccine)
	req := createRequest("GET", PublicURLV2+"/appointment/sessions/public/calendarByCenter", q, nil)
	req.Header.Set("User-Agent", "Cowin_Cli/1.0")
	req.Header.Set("Accept-Language", "hi_IN")
	err = sendRequest(req, &s)
	return s, err
}
