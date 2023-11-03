package files

import "github.com/spf13/cobra"

//var project string

// projectCmd represents the project command
var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func MakeFilesCmd() *cobra.Command {
	//filesCmd.AddCommand(makeGetCmd())
	filesCmd.AddCommand(uploadCmd)
	filesCmd.AddCommand(listCmd)
	filesCmd.AddCommand(deleteCmd)
	//filesCmd.AddCommand(GetCmd)
	//filesCmd.AddCommand(DeleteCmd)
	//filesCmd.PersistentFlags().StringVarP(&project, "project", "p", "", "Project apply files to")

	return filesCmd
}
