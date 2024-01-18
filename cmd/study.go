package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vanceji/shparentcredits/api"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	includeRequired  bool
	intervalSeconds  int32
	studentWhitelist string
	studentBlacklist string
	studyCmd         = &cobra.Command{
		Use:   "study",
		Short: "Visit courses",
		Long: `Visit courses to get credits. For example:

shparentcredits study -a https://www.example.com -t the-access-token [-s 123] [-c 456] [-r=false] `,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			api.Domain = host
			api.Token = token

			topModes := api.GetTopLevelMode()
			students := api.GetStudents(classId)
			//sort.SliceStable(students.Data, func(i, j int) bool {
			//	return students.Data[i].Points > students.Data[j].Points
			//})
			var wg sync.WaitGroup
			for _, student := range students.Data {
				wg.Add(1)
				go study(&wg, topModes, student, includeRequired)
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
	studyCmd.Flags().BoolVarP(&includeRequired, "include-required", "r", false, "whether to study the required courses, default false")
	studyCmd.Flags().Int32VarP(&intervalSeconds, "interval-seconds", "i", 30, "the time interval(random seconds) between two courses, default 30s")
	studyCmd.Flags().StringVarP(&studentWhitelist, "whitelist", "w", "", "the student name whitelist")
	studyCmd.Flags().StringVarP(&studentBlacklist, "blacklist", "b", "", "the student name blacklist")
	_ = studyCmd.MarkFlagRequired("host")
	_ = studyCmd.MarkFlagRequired("token")

	rootCmd.AddCommand(studyCmd)
}

func study(wg *sync.WaitGroup, topModes *api.Response[api.Mode], student api.Student, includeRequired bool) {
	defer wg.Done()
	now := time.Now().Local()

	if includeRequired {
		requiredCourseMode := api.GetSecondLevelModeForRequiredCourses(student.UserGuid)
		for _, mode := range requiredCourseMode.Data {
			requiredCourses := api.GetCourse(mode.ParentModeId, mode.ModeId, student.UserGuid)
			for _, requiredCourse := range requiredCourses.Data {
				time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				api.VisitCourse(student.UserGuid, requiredCourse.Id)
				fmt.Println(student.Username + " visited the required course /" +
					mode.ModeName + "/" + requiredCourse.Name + "," +
					time.Now().Format("2006/01/02 15:04:05"))
			}
		}
	}

	var MaxPoints = 600
	var points, _ = strconv.Atoi(student.Points)
	if points >= MaxPoints {
		fmt.Println(student.Username + "," + student.Points + ",points up to limited [" + student.Points + "]")
		return
	}
	courseCountToVisit := api.HowManyCoursesToVisitToday(student.UserGuid, now.Year(), int(now.Month()))
	//if points > 594 {
	//	courseCountToVisit = (MaxPoints - points) / 2
	//}
	fmt.Println(student.Username + "," + student.Points + ",need to visit [" + strconv.Itoa(courseCountToVisit) + "] courses.")
	if courseCountToVisit == 0 {
		return
	}

	if strings.TrimSpace(studentWhitelist) != "" && !isInlist(studentWhitelist, student.Username) {
		return
	}

	if strings.TrimSpace(studentBlacklist) != "" && isInlist(studentBlacklist, student.Username) {
		return
	}

	for _, parentMode := range topModes.Data {
		childModes := api.GetSecondLevelMode(parentMode.Id)
		for _, childMode := range childModes.Data {
			courses := api.GetCourse(parentMode.Id, childMode.Id, student.UserGuid)
			for _, course := range courses.Data {
				if course.IsVisited != "1" {
					time.Sleep(time.Duration(rand.Intn(int(intervalSeconds))+90) * time.Second)
					api.VisitCourse(student.UserGuid, course.Id)
					fmt.Println("\t" + student.Username + " visited the course '/" +
						parentMode.Name + "/" + childMode.Name + "/" + course.Name + "'," +
						time.Now().Format("2006/01/02 15:04:05"))
					courseCountToVisit--
				}
				if courseCountToVisit <= 0 {
					return
				}
			}
		}
	}
}

func isInlist(listStr string, name string) bool {
	var list = strings.Split(listStr, ",")
	for _, str := range list {
		if str == name {
			return true
		}
	}
	return false
}
