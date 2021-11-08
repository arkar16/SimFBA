package structs

type BaseTeam struct {
	TeamName    string
	Mascot      string
	TeamAbbr    string
	Coach       string
	City        string
	State       string
	Country     string
	Enrollment  int
	FirstPlayed int
	ColorOne    string
	ColorTwo    string
}

func (bt *BaseTeam) RemoveUserFromTeam() {
	bt.Coach = ""
}
