package api

import (
	"net/url"
)

type Session struct {
	CenterName        string   `json:"name"`
	Slots             []string `json:"slots"`
	Address           string   `json:"address"`
	District          string   `json:"district_name"`
	Block             string   `json:"block_name"`
	From              string   `json:"from"`
	To                string   `json:"to"`
	Fee               string   `json:"fee"`
	MinAge            int      `json:"min_age_limit"`
	Vaccine           string   `json:"string"`
	AvailableCapacity int      `json:"available_capacity"`
	AvailableDose1    int      `json:"available_capacity_dose1"`
	AvailableDose2    int      `json:"available_capacity_dose2"`
}

type AllSessions struct {
	Sessions []Session `json:"sessions"`
}

func GetAppointmentSessionsByPin(pincode string, appointmentDate string) (s AllSessions, err error) {
	q := url.Values{}
	q.Add("pincode", pincode)
	q.Add("date", appointmentDate)
	req := createRequest("GET", PublicURLV2+"/appointment/sessions/public/findByPin", q, nil)
	req.Header.Set("User-Agent", "Cowin_Cli/1.0")
	req.Header.Set("Accept-Language", "hi_IN")
	err = sendRequest(req, &s)
	return s, err
}
