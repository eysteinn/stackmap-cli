/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package project

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/eysteinn/stackmap-cli/pkg/apirequest"
	"github.com/eysteinn/stackmap-cli/pkg/global"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		project := args[0]

		data := url.Values{
			"name": {project},
		}

		endpoint, err := url.JoinPath(global.DefaultStackmapUrl, "projects")
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("Endpoint:", endpoint)
		resp, err := http.PostForm(endpoint, data)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("RespCode: ", resp.StatusCode)

		defer resp.Body.Close()

		/*var res map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&res)
		fmt.Println(res)
		fmt.Println(res["form"])
		if err != nil {
			log.Fatal(err)
		}*/

		respdata, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(string(respdata))
		apiresp := apirequest.ApiResponseSuccess{}
		err = json.Unmarshal(respdata, &apiresp)
		if err != nil {
			log.Fatal(err)
		}

		err = apiresp.GetError()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Project created successfully.")
	},
}

func init() {
	//project.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
