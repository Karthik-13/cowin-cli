package cmd

/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Karthik-13/cowin-cli/api"
	"github.com/spf13/cobra"
)

var (
	pincode string
	date    string
)

// appointmentCmd represents the appointment command
func getAppointmentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "appointments",
		Short: "Get info about Co-Win Vacccination appointment sessions and centers",
		Long:  `Get Co-Win Vaccination appointment sessions and centers available in India by district, PIN, Latitude and Longitude, Date`,
		Run: func(cmd *cobra.Command, args []string) {
			var data [][]string

			res, err := api.GetAppointmentSessionsByPin(pincode, date)
			if err != nil {
				fmt.Printf("Couldn't get appointments info - %v\n", err)
				return
			}

			if len(res.Sessions) == 0 {
				fmt.Println("Couldn't find any sessions for given date and Pincode")
				return
			}

			for _, session := range res.Sessions {
				row := []string{
					session.CenterName,
					session.Address,
					strings.Join(session.Slots, ","),
					session.District,
					session.From,
					session.To,
					session.Fee,
					strconv.Itoa(session.MinAge),
					session.Vaccine,
					strconv.Itoa(session.AvailableCapacity),
					strconv.Itoa(session.AvailableDose1),
					strconv.Itoa(session.AvailableDose2),
				}
				data = append(data, row)
			}

			table := getTable()

			table.SetHeader([]string{"Center", "Address", "Slots", "District", "From", "To", "Fee", "MinimumAge", "Vaccine", "AvailableCapacity", "Dose1", "Dose2"})

			for _, v := range data {
				table.Append(v)
			}

			table.Render()
		},
	}
	cmd.Flags().StringVar(&pincode, "pin-code", "", "Get appointment sesssions available near the pincode")
	cmd.Flags().StringVar(&date, "date", "", "Date for which we need the appointments. Format - DD-MM-YYYY")
	cmd.MarkFlagRequired("date")
	cmd.MarkFlagRequired("pin-code")

	return cmd
}

type AppointmentCenter struct {
	ID        string    `json:"center_id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	District  string    `json:"district_name"`
	State     string    `json:"state_name"`
	Block     string    `json:"block_name"`
	Pincode   string    `json:"pincode"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
	FeeType   string    `json:"fee_type"`
	Latitude  float64   `json:"lat"`
	Longitude float64   `json:"long"`
}

// func getAppointmentByLatLong(lat float64, long float64) []AppointmentCenter {
// 	res, err := http.Get(
// 		cowinApi + "appointment/centers/public/findByLatLong?lat=" + strconv.FormatFloat(lat, 'f', 5, 64) + "&long=" + strconv.FormatFloat(long, 'f', 5, 64),
// 	)

// 	if err != nil {
// 		log.Printf(err.Error())
// 	}
// 	json.Unmarshal(res)

// 	return
// }
