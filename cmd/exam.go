package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vanceji/shparentcredits/api"
	"strconv"
)

var (
	examCmd = &cobra.Command{
		Use:   "exam",
		Short: "Have a test",
		Long: `For example:

shparentcredits exam -a https://www.example.com -t the-access-token [-s 123] [-c 456]`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			api.Domain = host
			api.Token = token

			students := api.GetStudents(classId)
			for _, student := range students.Data {
				examSecondaryModes := api.GetSecondLevelModeForExam(student.UserGuid)
				for _, secondaryMode := range examSecondaryModes.Data {
					papers := api.GetPapers(student.UserGuid, secondaryMode.ModeId)
					for _, paper := range papers.Data {
						questions := api.GetQuestions(student.UserGuid, paper.PaperId)
						fmt.Println(student.Username + " had a test for paper [" +
							paper.PaperName + "], questions count: " + strconv.Itoa(len(questions.Data)))
						for _, q := range questions.Data {
							api.Answer(int(schoolId), student.UserGuid, paper.PaperId, q.QuestionId, q.CorrectAnswer)
						}
					}
				}
			}
		},
	}
)

func init() {
	examCmd.Flags().StringVarP(&host, "host", "a", "", "host of website")
	examCmd.Flags().Int32VarP(&schoolId, "school-id", "s", 101689, "school id, default 101689(think together)")
	examCmd.Flags().Int32VarP(&classId, "class-id", "c", 13989, "class id, default 13989(seeding-4)")
	examCmd.Flags().StringVarP(&token, "token", "t", "", "access token")

	_ = examCmd.MarkFlagRequired("host")
	_ = examCmd.MarkFlagRequired("token")

	rootCmd.AddCommand(examCmd)
}
