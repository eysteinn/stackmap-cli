package files

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"

	"github.com/eysteinn/stackmap-cli/pkg/apirequest"
	"github.com/eysteinn/stackmap-cli/pkg/global"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var uploadCmd = &cobra.Command{
	Use:   "upload [project] [product] [files...]",
	Short: "Upload a file to a project",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		project := args[0]
		product := args[1]
		/*endpoint, err := url.JoinPath(global.DefaultStackmapUrl, "projects", project, "files")
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Get(endpoint)
		if err != nil {
			log.Fatal(err)
		}
		//We Read the response body on the line below.
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))
		*/

		for _, arg := range args[2:] {
			files, err := filepath.Glob(arg)
			if err != nil {
				log.Fatal(err)
			}
			for _, file := range files {
				fmt.Println("Uploading file:", file)
				endpoint, err := url.JoinPath(global.DefaultStackmapUrl, "projects", project, "files")
				if err != nil {
					log.Fatal(err)
				}
				//resps := []apirequest.ApiResponseSuccess{}

				form := map[string]string{"image": "@" + file, "product": product, "project": project}
				err = apirequest.PostForm(endpoint, form, nil)
				if err != nil {
					log.Fatal(err)
				}
				/*for _, resp := range resps {
					if err = resp.GetError(); err != nil {
						log.Fatal(err)
					}
				}*/

				/*err := uploadFile(file, project, "hrit-ash")
				if err != nil {
					log.Fatal(err)
				}*/
			}
		}
		//project := args[0]
		//http.PostForm("stackmap.clouds.is:8080")
	},
}
