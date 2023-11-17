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
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint, err := url.JoinPath(global.DefaultStackmapUrl, "projects")
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Get(endpoint)
		if err != nil {
			log.Fatal(err)
		}
		//We Read the response body on the line below.
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		switch global.Output {
		case "json":
			fmt.Println(string(data))
		case "table":
			log.Fatal("Not implemented yet")
		default:
			apiresp := ListResponse{}
			err = json.Unmarshal(data, &apiresp)
			if err != nil {
				log.Fatal(err)
			}

			err = apiresp.GetError()
			if err != nil {
				log.Fatal(err)
			}

			for _, proj := range apiresp.Projects {
				fmt.Println(proj.Name)
			}
		}
		/*if global.Output == "table" {
			list := ListResponse{}
			err = json.Unmarshal(data, &list)
			if err != nil {
				log.Fatal(err)
			}
			values := make([]ApiResponseProject, 0, len(list.Projects))
			for _, v := range list.Projects {
				values = append(values, v)
			}
		}*/
		//project := args[0]
		//http.PostForm("stackmap.clouds.is:8080")
	},
}

type ApiResponseProject struct {
	Name string `json:"name,omitempty"`
	WMS  string `json:"wms,omitempty"`
}
type ListResponse struct {
	apirequest.ApiResponseSuccess
	ProjectNames []string                      `json:"project_names,omitempty"`
	Projects     map[string]ApiResponseProject `json:"projects,omitempty"`
}
