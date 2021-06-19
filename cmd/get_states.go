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
	"text/tabwriter"

	"github.com/Karthik-13/cowin-cli/api"
	"github.com/spf13/cobra"
)

const (
	ALIGN_DEFAULT = iota
	ALIGN_CENTER
	ALIGN_RIGHT
	ALIGN_LEFT
)

// getStatesCmd represents the getStates command
func (table *tabwriter.Ne) getStatesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "states",
		Short: "List of states in India",
		Long:  `Prints a list of all the states in India`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			var data [][]string

			res, err := api.GetStates()
			if err != nil {
				fmt.Printf("Couldn't get states info - %v", err)
				return
			}

			for _, state := range res.States {
				row := []string{strconv.Itoa(state.ID), state.Name}
				data = append(data, row)
			}

			table.SetHeader([]string{"ID", "Name"})

			for _, v := range data {
				table.Append(v)
			}

			table.Render()
		},
	}

	return cmd
}
