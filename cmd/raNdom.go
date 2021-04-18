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
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// raNdomCmd represents the raNdom command
var raNdomCmd = &cobra.Command{
	Use:   "raNdom",
	Short: "short description who cares",
	Long:  `gives random quotes of Lord supreme kanYE`,
	Run: func(cmd *cobra.Command, args []string) {
		getQuote()
	},
}

func init() {
	rootCmd.AddCommand(raNdomCmd)
}

type Quote struct {
	Kanye string `json:"quote"`
}

func getQuote() {
	apiUrl := "https://api.kanye.rest"
	quoteStore := Quote{}
	if err := json.Unmarshal(quoteData(apiUrl), &quoteStore); err != nil {
		log.Fatal("unmarshalling failed", err)
	}
	//fmt.Println("the type of the unmarshelled data: ", reflect.TypeOf(quoteStore.Kanye))
	fmt.Println("ye:", quoteStore.Kanye)
}

func quoteData(api string) []byte {
	req, err := http.NewRequest(
		http.MethodGet,
		api,
		nil,
	)
	if err != nil {
		log.Fatal("Unable to process the request", err)
	}
	req.Header.Add("Accept", "json/string")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Response not obtained", err)
	}
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("the type of the response body: ", reflect.TypeOf(responseBody))
	return responseBody

}
