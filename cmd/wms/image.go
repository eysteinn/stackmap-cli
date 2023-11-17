package wms

import (
	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:   "image [project] [uuid...]",
	Short: "Create image from WMS service",
	//Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		/*project
		apirequest.Get(global.DefaultStackmapHost + "/services/wms?")*/
	},
}
