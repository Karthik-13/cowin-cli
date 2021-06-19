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
	"strconv"

	"github.com/Karthik-13/cowin-cli/api"
	"github.com/spf13/cobra"
)

var stateId string

// getDistrictsCmd represents the districts command
func getDistrictsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "districts",
		Short: "List of districts in given state id",
		Long:  `Prints a list of all the districts in given state id`,
		Run: func(cmd *cobra.Command, args []string) {
			var data [][]string

			res, err := api.GetDistricts(stateId)
			if err != nil {
				fmt.Printf("Couldn't get districts info - %v", err)
				return
			}

			for _, district := range res.Districts {
				row := []string{strconv.Itoa(district.ID), district.Name}
				data = append(data, row)
			}

			table := getTable()

			table.SetHeader([]string{"ID", "Name"})

			for _, v := range data {
				table.Append(v)
			}

			table.Render()
		},
	}

	cmd.Flags().StringVar(&stateId, "state-id", "", "Specify a state id")
	cmd.MarkFlagRequired("state-id")

	return cmd
}
