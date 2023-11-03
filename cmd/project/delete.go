/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package project

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/eysteinn/stackmap-cli/pkg/global"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			err := DeleteProject(arg)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func DeleteProject(name string) error {
	endpoint, err := url.JoinPath(global.DefaultStackmapUrl, "projects", name)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	// create a new DELETE request
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// print the response body
	fmt.Println(string(body))
	return nil
	//project := args[0]
	//http.PostForm("stackmap.clouds.is:8080")
}
