package util

func GetRoundAbbreviation(str string) string {
	if str == "1" {
		return "1st Round"
	} else if str == "2" {
		return "2nd Round"
	} else if str == "3" {
		return "3rd Round"
	} else if str == "4" {
		return "4th Round"
	} else if str == "5" {
		return "5th Round"
	} else if str == "6" {
		return "6th Round"
	}
	return "7th Round"
}
