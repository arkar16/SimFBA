package structs

type BaseTeam struct {
	TeamName         string
	Mascot           string
	TeamAbbr         string
	Coach            string
	City             string
	State            string
	Country          string
	StadiumID        uint
	Stadium          string
	StadiumCapacity  int
	RecordAttendance int
	Enrollment       int
	FirstPlayed      int
	ColorOne         string
	ColorTwo         string
	ColorThree       string
	DiscordID        string
}

func (bt *BaseTeam) RemoveUserFromTeam() {
	bt.Coach = "AI"
}

func (bt *BaseTeam) AssignUserToTeam(user string) {
	bt.Coach = user
}

func (bt *BaseTeam) AssignDiscordID(id string) {
	bt.DiscordID = id
}
