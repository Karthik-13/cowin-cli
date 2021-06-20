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
	"log"
	"strconv"
	"strings"

	"github.com/Karthik-13/cowin-cli/api"
)

type AllVaccineCentersByPincode struct {
	Centers json.RawMessage `json:"centers"`
}

type VaccineCenters struct {
	VaccineCenterId       int    `json:"center_id"`
	VaccineCenterName     string `json:"name"`
	VaccineCenterState    string `json:"state_name"`
	VaccineCenterDistrict string `json:"district_name"`
	VaccineCenterBlock    string `json:"block_name"`
	VaccineCenterPincode  int    `json:"pincode"`
	VaccineFeeType        string `json:"fee_type"`
	//VaccineSessions       json.RawMessage `json:"sessions"`
	VaccineSessions []VaccineSession `json:"sessions"`
}

type VaccineSession struct {
	VaccinationDate             string   `json:"date"`
	VaccinationAvailableCapcity int      `json:"available_capacity"`
	VaccinationAgeLimit         int      `json:"min_age_limit"`
	VaccineName                 string   `json:"vaccine"`
	VaccineFirstDose            int      `json:"available_capacity_dose1"`
	VaccineSecondDose           int      `json:"available_capacity_dose2"`
	VaccineAvailableSlots       []string `json:"slots"`
}

// get the vaccination availability details based on the pincode / district
func getDataByPincodeDistrict(urlEndpoint string) {
	responseBytes := api.GetApiData(urlEndpoint)
	var data [][]string

	allVaccineCentersByPincodeRaw := AllVaccineCentersByPincode{}

	if err := json.Unmarshal(responseBytes, &allVaccineCentersByPincodeRaw); err != nil {
		log.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	vaccineCenters := []VaccineCenters{}

	if err := json.Unmarshal(allVaccineCentersByPincodeRaw.Centers, &vaccineCenters); err != nil {
		log.Printf("Could not unmarshal vaccine centers. %v", err)
	}
	table := api.GenerateTable()

	table.SetHeader([]string{"Vaccine Center ID", "Vaccine Center Name", "State", "District", "Block", "Pincode", "Vaccine", "Fee Type", "Vaccine Date", "Slot Timing", "Age Limit", "Vaccine Available", "Available First Dose", "Available Second Dose"})
	for _, v := range vaccineCenters {
		for _, v1 := range v.VaccineSessions {

			row := []string{strconv.Itoa(v.VaccineCenterId), v.VaccineCenterName, v.VaccineCenterState, v.VaccineCenterDistrict, v.VaccineCenterBlock, strconv.Itoa(v.VaccineCenterPincode), v1.VaccineName, v.VaccineFeeType, v1.VaccinationDate, strings.Join(v1.VaccineAvailableSlots, ",\n"), strconv.Itoa(v1.VaccinationAgeLimit) + "+", strconv.Itoa(v1.VaccinationAvailableCapcity), strconv.Itoa(v1.VaccineFirstDose), strconv.Itoa(v1.VaccineSecondDose)}
			data = append(data, row)

		}
	}
	table.AppendBulk(data)
	table.Render()
}

// get the vaccination availability details based on the vaccination center
func getDataByCenter(urlEndpoint string) {
	responseBytes := api.GetApiData(urlEndpoint)
	var data [][]string
	// var vaccinationCenterData map[string]interface{}

	allVaccineCentersByPincodeRaw := AllVaccineCentersByPincode{}

	if err := json.Unmarshal(responseBytes, &allVaccineCentersByPincodeRaw); err != nil {
		log.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	v := VaccineCenters{}

	if err := json.Unmarshal(allVaccineCentersByPincodeRaw.Centers, &v); err != nil {
		log.Printf("Could not unmarshal vaccine centers. %v", err)
	}
	table := api.GenerateTable()

	table.SetHeader([]string{"Vaccine Center ID", "Vaccine Center Name", "State", "District", "Block", "Pincode", "Vaccine", "Fee Type", "Vaccine Date", "Slot Timing", "Age Limit", "Vaccine Available", "Available First Dose", "Available Second Dose"})

	for _, v1 := range v.VaccineSessions {

		row := []string{strconv.Itoa(v.VaccineCenterId), v.VaccineCenterName, v.VaccineCenterState, v.VaccineCenterDistrict, v.VaccineCenterBlock, strconv.Itoa(v.VaccineCenterPincode), v1.VaccineName, v.VaccineFeeType, v1.VaccinationDate, strings.Join(v1.VaccineAvailableSlots, ",\n"), strconv.Itoa(v1.VaccinationAgeLimit) + "+", strconv.Itoa(v1.VaccinationAvailableCapcity), strconv.Itoa(v1.VaccineFirstDose), strconv.Itoa(v1.VaccineSecondDose)}
		data = append(data, row)

	}

	table.AppendBulk(data)
	table.Render()
}
