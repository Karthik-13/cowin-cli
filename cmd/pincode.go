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
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Karthik-13/cowin-cli/api"
	"github.com/spf13/cobra"
)

// pincodeCmd represents the pincode command
var searchDate, vaccine string

var pincodeCmd = &cobra.Command{
	Use:   "pincode",
	Short: "Gets the vaccination calendar using pincode",
	Long:  `Access CoWin API and gets the vaccination calendar using the pincode,prints the hospitals and vaccine available at the place`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Requires a district pincode as argument")
		} else if len(args) > 1 {
			return errors.New("Too many argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		pincode := args[0]
		getDataByPincode(pincode, searchDate, vaccine)
	},
}

func init() {
	findbyCmd.AddCommand(pincodeCmd)
	d := time.Now()
	date := fmt.Sprintf("%d-%02d-%02d", d.Day(), int(d.Month()), d.Year())
	pincodeCmd.PersistentFlags().StringVar(&searchDate, "date", date, "Pass \"date\" in format (dd-mm-yyyy) to get the list of vaccination centers. Defaults to current date.")
	pincodeCmd.PersistentFlags().StringVar(&vaccine, "vaccine", "", "Vaccine name to customize the search	")
}

type AllVaccineCentersByPincode struct {
	Centers json.RawMessage `json:"centers"`
}

type VaccineCenters struct {
	VaccineCenterName     string `json:"name"`
	VaccineCenterState    string `json:"state_name"`
	VaccineCenterDistrict string `json:"district_name"`
	VaccineCenterBlock    string `json:"block_name"`
	VaccineCenterPincode  int    `json:"pincode"`
	//VaccineSessions       json.RawMessage `json:"sessions"`
	VaccineSessions []VaccineSession `json:"sessions"`
}

type VaccineSession struct {
	VaccinationDate             string `json:"date"`
	VaccinationAvailableCapcity int    `json:"available_capacity"`
	VaccinationAgeLimit         int    `json:"min_age_limit"`
	VaccineName                 string `json:"vaccine"`
	VaccineFirstDose            int    `json:"available_capacity_dose1"`
	VaccineSecondDose           int    `json:"available_capacity_dose2"`
}

func getDataByPincode(pincode string, date string, vaccine string) {
	baseUrl := fmt.Sprintf("https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByPin?pincode=%s&date=%s&vaccine=%s", pincode, date, vaccine)
	responseBytes := api.GetApiData(baseUrl)
	var data [][]string
	// var vaccinationCenterData map[string]interface{}

	allVaccineCentersByPincodeRaw := AllVaccineCentersByPincode{}

	if err := json.Unmarshal(responseBytes, &allVaccineCentersByPincodeRaw); err != nil {
		log.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	vaccineCenters := []VaccineCenters{}

	if err := json.Unmarshal(allVaccineCentersByPincodeRaw.Centers, &vaccineCenters); err != nil {
		log.Printf("Could not unmarshal vaccine centers. %v", err)
	}
	table := api.GenerateTable()

	table.SetHeader([]string{"Vaccine Center Name", "State", "District", "Block", "Pincode", "Vaccine", "Vaccine Date", "Age Limit", "Vaccine Available", "Available First Dose", "Available Second Dose"})
	for _, v := range vaccineCenters {
		for _, v1 := range v.VaccineSessions {
			row := []string{v.VaccineCenterName, v.VaccineCenterState, v.VaccineCenterDistrict, v.VaccineCenterBlock, strconv.Itoa(v.VaccineCenterPincode), v1.VaccineName, v1.VaccinationDate, strconv.Itoa(v1.VaccinationAgeLimit), strconv.Itoa(v1.VaccinationAvailableCapcity), strconv.Itoa(v1.VaccineFirstDose), strconv.Itoa(v1.VaccineSecondDose)}
			data = append(data, row)
		}
	}
	table.AppendBulk(data)
	table.Render()
}
