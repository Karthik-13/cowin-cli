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

	"github.com/spf13/cobra"
)

// pincodeCmd represents the pincode command
var pincodeCmd = &cobra.Command{
	Use:   "pincode",
	Short: "Gets the vaccination calendar using pincode",
	Long:  `Access CoWin API and gets the vaccination calendar using the pincode,prints the hospitals and vaccine available at the place`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pincode called")
	},
}

func init() {
	findbyCmd.AddCommand(pincodeCmd)
}
