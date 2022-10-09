package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vanceji/shparentcredits/api"
)

var (
	reportCmd = &cobra.Command{
		Use:   "report",
		Short: "Credits Report",
		Long: `Show credits details for current school. For example:

shparentcredits report -a https://www.example.com -t the-access-token [-s 123] [-c 456]`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			api.Domain = host
			api.Token = token

			api.PrintClassRanking(int(schoolId))
		},
	}
)

func init() {
	reportCmd.Flags().StringVarP(&host, "host", "a", "", "host of website")
	reportCmd.Flags().Int32VarP(&schoolId, "school-id", "s", 101689, "school id, default 101689(think together)")
	reportCmd.Flags().Int32VarP(&classId, "class-id", "c", 13989, "class id, default 13989(seeding-4)")
	reportCmd.Flags().StringVarP(&token, "token", "t", "", "access token")

	_ = reportCmd.MarkFlagRequired("host")
	_ = reportCmd.MarkFlagRequired("token")

	rootCmd.AddCommand(reportCmd)
}
