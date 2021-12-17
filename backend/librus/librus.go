package librus

import (
	"encoding/json"
	"net/url"
	"os"
)

type Librus struct {
	AccessToken string `json:"access_token"`
	Subjects    []Subject
	Lessons     []Lesson
	Attendances []Attendance
}

type CalculatedNode struct {
	Name       string
	Attendance [2][5]int
}

type App interface {
	LoadEndpoints(endpoints ...string) error
	MakeLessonSubjectMap() map[int]*CalculatedNode
	GetAttendance(map[int]*CalculatedNode)
}

var API_ROOT string = os.Getenv("API_ROOT")
var API_TOKEN string = os.Getenv("API_TOKEN")

func Init(username, password string) (App, error) {
	user := url.Values{}
	user.Set("username", username)
	user.Set("password", password)
	user.Set("grant_type", "password")
	user.Set("librus_long_term_token", "0")
	user.Set("librus_rules_accepted", "true")
	user.Set("librus_mobile_rules_accepted", "true")

	response, err := request(&mRequestParams{
		API_ROOT + "OAuth/Token",
		API_TOKEN,
		"POST",
		user.Encode(),
	})
	if err != nil {
		return nil, err
	}

	librus := &Librus{}
	json.Unmarshal(response, librus)

	return librus, nil
}
