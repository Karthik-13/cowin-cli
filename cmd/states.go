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
	"os"
	"strconv"

	"github.com/Karthik-13/cowin-cli/api"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// statesCmd represents the states command
var statesCmd = &cobra.Command{
	Use:   "states",
	Short: "List of states available in India.",
	Long:  `Access COWIN public api to get the list of states available in India.`,
	Run: func(cmd *cobra.Command, args []string) {
		getStates()
	},
}

func init() {
	getCmd.AddCommand(statesCmd)
}

type StatesList struct {
	States json.RawMessage `json:"states"`
}

type States struct {
	StateID   int    `json:"state_id"`
	StateName string `json:"state_name"`
}

func getStates() {
	baseUrl := "https://cdn-api.co-vin.in/api/v2/admin/location/states"
	responseBytes := api.GetApiData(baseUrl)

	statesListRaw := StatesList{}

	if err := json.Unmarshal(responseBytes, &statesListRaw); err != nil {
		log.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	statesList := []States{}

	if err := json.Unmarshal(statesListRaw.States, &statesList); err != nil {
		log.Printf("Could not unmarshal states list. %v", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"State Id", "State Name"})
	for _, element := range statesList {
		stateId := strconv.Itoa(element.StateID)
		data := []string{stateId, element.StateName}
		table.Append(data)
	}
	table.Render()
}
