package files

import (
	"fmt"
	"log"
	"net/url"

	"github.com/eysteinn/stackmap-cli/pkg/apirequest"
	"github.com/eysteinn/stackmap-cli/pkg/global"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [project] [uuid...]",
	Short: "Delete files in a project",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		project := args[0]

		for _, uuid := range args[1:] {
			endpoint, err := url.JoinPath(global.DefaultStackmapUrl, "projects", project, "files", uuid)
			if err != nil {
				log.Fatal(err)
			}
			if !global.Quiet {
				fmt.Println("Deleting:", uuid)
			}
			err = apirequest.Delete(endpoint)
			if err != nil {
				log.Fatal(err)
			}
		}

	},
}
