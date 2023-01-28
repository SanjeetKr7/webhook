package pkg

type SchoolDetail struct {
	SchoolName string        `json:"schoolName"`
	SchoolId   int           `json:"schoolId"`
	WebHook    WebHookDetail `json:"web"`
}

type WebHookDetail struct {
	EventName   string `json:"eventname"`
	EndPointUrl string `json:"endpoint"`
}

type StudentDetail struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	SchoolId int    `json:"schoolId"`
}
