package structs

type BaseTeam struct {
	TeamName         string
	Mascot           string
	TeamAbbr         string
	Coach            string
	City             string
	State            string
	Country          string
	Stadium          string
	StadiumCapacity  int
	RecordAttendance int
	Enrollment       int
	FirstPlayed      int
	ColorOne         string
	ColorTwo         string
	ColorThree       string
}

func (bt *BaseTeam) RemoveUserFromTeam() {
	bt.Coach = "AI"
}

func (bt *BaseTeam) AssignUserToTeam(user string) {
	bt.Coach = user
}
