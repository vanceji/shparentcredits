package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vanceji/shparentcredits/api"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	studyCmd = &cobra.Command{
		Use:   "study",
		Short: "Visit courses",
		Long: `Visit courses to get credits. For example:

shparentcredits study -a https://www.example.com -t the-access-token [-s 123] [-c 456] `,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			api.Domain = host
			api.Token = token

			topModes := api.GetTopLevelMode()
			students := api.GetStudents(classId)
			var wg sync.WaitGroup
			for _, student := range students.Data {
				wg.Add(1)
				go study(&wg, topModes, student)
			}
			wg.Wait()
		},
	}
)

func init() {
	studyCmd.Flags().StringVarP(&host, "host", "a", "", "host of website")
	studyCmd.Flags().Int32VarP(&schoolId, "school-id", "s", 101689, "school id, default 101689(think together)")
	studyCmd.Flags().Int32VarP(&classId, "class-id", "c", 13989, "class id, default 13989(seeding-4)")
	studyCmd.Flags().StringVarP(&token, "token", "t", "", "access token")

	_ = studyCmd.MarkFlagRequired("host")
	_ = studyCmd.MarkFlagRequired("token")

	rootCmd.AddCommand(studyCmd)
}

func study(wg *sync.WaitGroup, topModes *api.Response[api.Mode], student api.Student) {
	defer wg.Done()
	now := time.Now().Local()
	courseCountToVisit := api.HowManyCoursesToVisitToday(student.UserGuid, now.Year(), int(now.Month()))
	fmt.Println(student.Username + " need to visit [" + strconv.Itoa(courseCountToVisit) + "] courses.")
	if courseCountToVisit == 0 {
		return
	}

	for _, parentMode := range topModes.Data {
		childModes := api.GetSecondLevelMode(parentMode.Id)
		for _, childMode := range childModes.Data {
			courses := api.GetCourse(parentMode.Id, childMode.Id, student.UserGuid)
			for _, course := range courses.Data {
				if course.IsVisited != "1" {
					time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
					api.VisitCourse(student.UserGuid, course.Id)
					fmt.Println("\t" + student.Username + " visited the course '/" + parentMode.Name + "/" + childMode.Name + "/" + course.Name + "'," + time.Now().Format("2006/01/02 15:04:05"))
					courseCountToVisit--
				}
				if courseCountToVisit <= 0 {
					return
				}
			}
		}
	}
}
