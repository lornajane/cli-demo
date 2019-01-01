// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
)

type data struct {
	Count   int `json:"count"`
	Numbers []struct {
		Msisdn  string `json:"msisdn"`
		Country string `json:"country"`
	} `json:"numbers"`
}

// numbersCmd represents the numbers command
var numbersCmd = &cobra.Command{
	Use:        "numbers",
	Aliases:    []string{"number"},
	SuggestFor: []string{"list", "rhubarb"},
	Short:      "Work with your Nexmo numbers",
	Long:       `This command and its subcommands enable you to find, buy, list, configure and cancel your Nexmo telephone numbers.`,
	Run: func(cmd *cobra.Command, args []string) {
		// prepare the URL (auth in query string)
		url := fmt.Sprint(
			"https://rest.nexmo.com/account/numbers?api_key=",
			viper.Get("api-key"),
			"&api_secret=",
			viper.Get("api-secret"),
		)
		fmt.Println(url)

		// make API call
		client := http.Client{}
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		response, getErr := client.Do(request)
		if getErr != nil {
			log.Fatal(getErr)
		}

		body, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		data1 := data{}
		jsonErr := json.Unmarshal(body, &data1)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		for i := 0; i < len(data1.Numbers); i++ {
			fmt.Println(data1.Numbers[i].Msisdn)
		}
	},
}

func init() {
	rootCmd.AddCommand(numbersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// numbersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// numbersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
