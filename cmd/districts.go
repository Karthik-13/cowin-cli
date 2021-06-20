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
	"fmt"
	"log"
	"strconv"

	"github.com/Karthik-13/cowin-cli/api"
	"github.com/spf13/cobra"
)

// districtsCmd represents the districts command
var districtsCmd = &cobra.Command{
	Use:   "districts",
	Short: "List of districts available in the state.",
	Long:  `Access COWIN public api to get the list of districts available in the state.`,
	Run: func(cmd *cobra.Command, args []string) {
		state_id, _ := cmd.Flags().GetString("state_id")
		getDistricts(state_id)
	},
}

var stateId string

type DistrictList struct {
	Districts json.RawMessage `json:"districts"`
}

type Districts struct {
	DistrictId   int    `json:"district_id"`
	DistrictName string `json:"district_name"`
}

func init() {
	getCmd.AddCommand(districtsCmd)
	districtsCmd.PersistentFlags().StringVar(&stateId, "state_id", "", "Pass \"state_id\" to get the list of districts available for the state.\nGet state list using \"cowin get states\"")
	districtsCmd.MarkPersistentFlagRequired("state_id")

}

func getDistricts(stateId string) {
	baseUrl := fmt.Sprintf("https://cdn-api.co-vin.in/api/v2/admin/location/districts/%s", stateId)
	responseBytes := api.GetApiData(baseUrl)

	districtListRaw := DistrictList{}

	if err := json.Unmarshal(responseBytes, &districtListRaw); err != nil {
		log.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	districtList := []Districts{}

	if err := json.Unmarshal(districtListRaw.Districts, &districtList); err != nil {
		log.Printf("Could not unmarshal states list. %v", err)
	}

	table := api.GenerateTable()
	table.SetHeader([]string{"District Id", "District Name"})
	for _, element := range districtList {
		data := []string{strconv.Itoa(element.DistrictId), element.DistrictName}
		table.Append(data)
	}
	table.Render()
}
