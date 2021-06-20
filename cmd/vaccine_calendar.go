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
package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

const (
	PublicURLV2 = "https://cdn-api.co-vin.in/api/v2"
)

var searchDate, vaccine, pincode, districtId, centerId string
var apiEndpoint string

// vaccineCalendarCmd represents the vaccineCalendar command
func getVaccineCalendarCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "vaccine-calendar",
		Short: "Get vaccination sessions for 7-days.",
		Long:  `Get the vaccination slot available for 7 days by initiating search using pincode, district or vaccine center.`,
		Run: func(cmd *cobra.Command, args []string) {
			if pincode == "" && districtId == "" && centerId == "" {
				log.Error("pincode or district-id or center-id is mandatory")
			}

			if pincode != "" && districtId == "" && centerId == "" {
				apiEndpoint = PublicURLV2 + "/appointment/sessions/public/calendarByPin?pincode=" + pincode + "&date=" + searchDate + "&vaccine=" + vaccine
				getDataByPincodeDistrict(apiEndpoint)
			}

			if districtId != "" && pincode == "" && centerId == "" {
				apiEndpoint = PublicURLV2 + "/appointment/sessions/public/calendarByDistrict?district_id=" + districtId + "&date=" + searchDate + "&vaccine=" + vaccine
				getDataByPincodeDistrict(apiEndpoint)
			}

			if centerId != "" && pincode == "" && districtId == "" {
				apiEndpoint = PublicURLV2 + "/appointment/sessions/public/calendarByCenter?center_id=" + centerId + "&date=" + searchDate + "&vaccine=" + vaccine
				getDataByCenter(apiEndpoint)
			}

		},
	}

	d := time.Now()
	date := fmt.Sprintf("%d-%02d-%02d", d.Day(), int(d.Month()), d.Year())

	cmd.PersistentFlags().StringVar(&pincode, "pincode", "", "Pincode of the district/state")
	cmd.PersistentFlags().StringVar(&districtId, "district-id", "", "Pass \"district-id\" to get the list vaccination centers and slots available for appointment.\nGet district list using \"cowin get district --state_id\"")
	cmd.PersistentFlags().StringVar(&centerId, "center-id", "", "Pass \"center-id\" to get the list vaccination centers and slots available for appointment.")
	cmd.PersistentFlags().StringVar(&searchDate, "date", date, "Pass \"date\" in format (dd-mm-yyyy) to get the list of vaccination centers. Defaults to current date.")
	cmd.PersistentFlags().StringVar(&vaccine, "vaccine", "", "Vaccine name to customize the search	")

	return cmd
}
