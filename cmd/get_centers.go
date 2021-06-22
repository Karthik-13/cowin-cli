// /*
// Copyright © 2021 NAME HERE <EMAIL ADDRESS>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//
package cmd

import (
	"fmt"
	"github.com/Karthik-13/cowin-cli/api"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"time"
)

var (
	c_pincode    string
	c_date       string
	c_districtid string
	c_centerid   string
	c_vaccine    string
)

// getCentersCmd represents the getCenters command
func getCentersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "centers",
		Short: "List the vaccination centers in India - 7 days detail will be listed.",
		Long:  `Prints the list of vaccination centers, along with their location, pin, block - 7 days detail will be listed.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			var data [][]string

			if c_pincode == "" && c_districtid == "" && c_centerid == "" {
				fmt.Println("pincode or district-id or center-id is mandatory")
				return
			}

			// get centers list based on the pincode
			if c_pincode != "" && c_districtid == "" && c_centerid == "" {
				res, err := api.GetCentersSessionsByPin(c_pincode, c_date, c_vaccine)
				if err != nil {
					fmt.Printf("Couldn't get appointments info - %v\n", err)
					return
				}

				if len(res.Centers) == 0 {
					fmt.Println("Couldn't find any centers for given date and Pincode")
					return
				}

				for _, center := range res.Centers {
					for _, session := range center.Sessions {
						row := []string{
							strconv.Itoa(center.CenterId),
							center.CenterName,
							center.CenterState,
							center.CenterDistrict,
							center.CenterBlock,
							strconv.Itoa(center.CenterPincode),
							session.VaccineName,
							center.FeeType,
							session.Date,
							strings.Join(session.AvailableSlots, ",\n"),
							strconv.Itoa(session.AgeLimit) + "+",
							strconv.Itoa(session.AvailableCapcity),
							strconv.Itoa(session.FirstDose),
							strconv.Itoa(session.SecondDose),
						}
						data = append(data, row)
					}
				}
			}

			// get centers list based on the district id
			if c_districtid != "" && c_pincode == "" && c_centerid == "" {
				res, err := api.GetCentersSessionsByDistrict(c_districtid, c_date, c_vaccine)
				if err != nil {
					fmt.Printf("Couldn't get appointments info - %v\n", err)
					return
				}

				if len(res.Centers) == 0 {
					fmt.Println("Couldn't find any centers for given date and Pincode")
					return
				}

				for _, center := range res.Centers {
					for _, session := range center.Sessions {
						row := []string{
							strconv.Itoa(center.CenterId),
							center.CenterName,
							center.CenterState,
							center.CenterDistrict,
							center.CenterBlock,
							strconv.Itoa(center.CenterPincode),
							session.VaccineName,
							center.FeeType,
							session.Date,
							strings.Join(session.AvailableSlots, ",\n"),
							strconv.Itoa(session.AgeLimit) + "+",
							strconv.Itoa(session.AvailableCapcity),
							strconv.Itoa(session.FirstDose),
							strconv.Itoa(session.SecondDose),
						}
						data = append(data, row)
					}
				}
			}

			// get centers list based on the center id
			if c_centerid != "" && c_pincode == "" && c_districtid == "" {
				res, err := api.GetCentersSessionsByCenter(c_centerid, c_date, c_vaccine)
				if err != nil {
					fmt.Printf("Couldn't get appointments info - %v\n", err)
					return
				}

				center := res.Centers
				for _, session := range center.Sessions {
					row := []string{
						strconv.Itoa(center.CenterId),
						center.CenterName,
						center.CenterState,
						center.CenterDistrict,
						center.CenterBlock,
						strconv.Itoa(center.CenterPincode),
						session.VaccineName,
						center.FeeType,
						session.Date,
						strings.Join(session.AvailableSlots, ",\n"),
						strconv.Itoa(session.AgeLimit) + "+",
						strconv.Itoa(session.AvailableCapcity),
						strconv.Itoa(session.FirstDose),
						strconv.Itoa(session.SecondDose),
					}
					data = append(data, row)
				}

			}

			table := getTable()
			table.SetHeader([]string{"Center ID", "Center Name", "State", "District", "Block", "Pincode", "Vaccine Name", "Fee Type", "Vaccine Date", "Slot Timing", "Age Limit", "Vaccine Available", "Available First Dose", "Available Second Dose"})
			for _, v := range data {
				table.Append(v)
			}

			table.Render()
		},
	}

	d := time.Now()
	date := fmt.Sprintf("%d-%02d-%02d", d.Day(), int(d.Month()), d.Year())

	cmd.Flags().StringVar(&c_pincode, "pincode", "", "Get appointment sesssions available near the pincode")
	cmd.Flags().StringVar(&c_date, "date", date, "Date for which we need the appointments. Format - DD-MM-YYYY")
	cmd.Flags().StringVar(&c_districtid, "district-id", "", "Pass \"district-id\" to get the list vaccination centers and slots available for appointment.\nGet district list using \"cowin get district --state-id\"")
	cmd.Flags().StringVar(&c_centerid, "center-id", "", "Pass \"center-id\" to get the list vaccination centers and slots available for appointment.")
	cmd.Flags().StringVar(&c_vaccine, "vaccine", "", "Vaccine name to customize the search.")

	return cmd
}

// func init() {
// 	getCmd.AddCommand(getCentersCmd)

// 	// Here you will define your flags and configuration settings.

// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// getCentersCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// getCentersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// 	getCentersCmd.Flags().StringVarP(&filterBy, "source", "s", "", "Source directory to read from")

// }
