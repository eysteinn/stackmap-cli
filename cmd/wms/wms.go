package wms

import "github.com/spf13/cobra"

var project string

var filesCmd = &cobra.Command{
	Use:   "wms",
	Short: "Collection of wms actions",
}

func MakeWMSCmd() *cobra.Command {

	return filesCmd
}
