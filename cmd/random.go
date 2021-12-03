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
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "This cli tool outputs a random joke",
	Long:  `The tool pulls a random joke and displays it as output`,
	Run: func(cmd *cobra.Command, args []string) {
		getAJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int64  `json:"status"`
}

func getAJoke() {
	api := "https://icanhazdadjoke.com/"

	jokeBytes := getJokeData(api)
	joke_obj := Joke{}

	if err := json.Unmarshal(jokeBytes, &joke_obj); err != nil {
		fmt.Printf("Json unmarshal failed due to %v", err)
	}

	fmt.Println(joke_obj.Joke)
}

func getJokeData(url string) []byte {
	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		fmt.Printf("Http request failed due to %s", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "DadJoke CLI Tutorial")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error response : %s", err)
	}

	resBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Error reading response body : %s", err)
	}

	return resBytes
}
