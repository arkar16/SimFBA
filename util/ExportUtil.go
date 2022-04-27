package util

func GetNFLOverallGrade(value int) string {
	if value > 63 {
		return "A+"
	} else if value > 61 {
		return "A"
	} else if value > 59 {
		return "A-"
	} else if value > 57 {
		return "B+"
	} else if value > 55 {
		return "B"
	} else if value > 53 {
		return "B-"
	} else if value > 51 {
		return "C+"
	} else if value > 49 {
		return "C"
	} else if value > 47 {
		return "C-"
	} else if value > 45 {
		return "D+"
	} else if value > 43 {
		return "D"
	} else if value > 41 {
		return "D-"
	}
	return "F"
}

func GetOverallGrade(value int) string {
	if value > 44 {
		return "A"
	} else if value > 34 {
		return "B"
	} else if value > 24 {
		return "C"
	} else if value > 14 {
		return "D"
	}
	return "F"
}

func GetLetterGrade(Attribute int, mean float32, stddev float32) string {
	if mean == 0 || stddev == 0 {
		return GetOverallGrade(Attribute)
	}

	val := float32(Attribute)
	dev := stddev * 2
	if val > mean+dev {
		return "A"
	}
	dev = stddev
	if val > mean+dev {
		return "B"
	}
	if val > mean {
		return "C"
	}
	dev = stddev * -1
	if val > mean+dev {
		return "D"
	}
	return "F"
}

func GetYearAndRedshirtStatus(year int, redshirt bool) (string, string) {
	status := ""
	if redshirt {
		status = "Redshirt"
	} else {
		status = ""
	}

	if year == 1 && !redshirt {
		return "Fr", status
	} else if year == 2 && redshirt {
		return "(Fr)", status
	} else if year == 2 && !redshirt {
		return "So", status
	} else if year == 3 && redshirt {
		return "(So)", status
	} else if year == 3 && !redshirt {
		return "Jr", status
	} else if year == 4 && redshirt {
		return "(Jr)", status
	} else if year == 4 && !redshirt {
		return "Sr", status
	} else if year == 5 && redshirt {
		return "(Sr)", status
	}
	return "Super Sr", status
}
