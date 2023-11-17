package files

/*
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
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		endpoint, err := url.JoinPath(global.DefaultStackmapUrl, "projects", project, "files")
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

		type Data struct {
			ApiResponseSuccess
			Success bool `json:"files,omitempty"`
		}

		//if files, found := body[]
		fmt.Println(string(body))

		//project := args[0]
		//http.PostForm("stackmap.clouds.is:8080")
	},
}

func makeGetCmd() *cobra.Command {
	cmd := getCmd
	cmd.Flags().StringVarP(&project, "project", "p", "", "Project apply files to")
	return cmd
}*/
