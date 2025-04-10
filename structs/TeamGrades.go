package structs

type TeamGrade struct {
	OffenseGradeNumber      float64
	DefenseGradeNumber      float64
	SpecialTeamsGradeNumber float64
	OverallGradeNumber      float64
	OffenseGradeLetter      string
	DefenseGradeLetter      string
	SpecialTeamsGradeLetter string
	OverallGradeLetter      string
}

func (tg *TeamGrade) SetOffenseGradeNumber(grade float64) {
	tg.OffenseGradeNumber = grade
}

func (tg *TeamGrade) SetDefenseGradeNumber(grade float64) {
	tg.DefenseGradeNumber = grade
}

func (tg *TeamGrade) SetSpecialTeamsGradeNumber(grade float64) {
	tg.SpecialTeamsGradeNumber = grade
}

func (tg *TeamGrade) SetOverallGradeNumber(grade float64) {
	tg.OverallGradeNumber = grade
}

func (tg *TeamGrade) SetOffenseGradeLetter(grade string) {
	tg.OffenseGradeLetter = grade
}

func (tg *TeamGrade) SetDefenseGradeLetter(grade string) {
	tg.DefenseGradeLetter = grade
}

func (tg *TeamGrade) SetSpecialTeamsGradeLetter(grade string) {
	tg.SpecialTeamsGradeLetter = grade
}

func (tg *TeamGrade) SetOverallGradeLetter(grade string) {
	tg.OverallGradeLetter = grade
}

func (tg *TeamGrade) GetOffenseGradeNumber() float64 {
	return tg.OffenseGradeNumber
}

func (tg *TeamGrade) GetDefenseGradeNumber() float64 {
	return tg.DefenseGradeNumber
}

func (tg *TeamGrade) GetSpecialTeamsGradeNumber() float64 {
	return tg.SpecialTeamsGradeNumber
}

func (tg *TeamGrade) GetOverallGradeNumber() float64 {
	return tg.OverallGradeNumber
}

func (tg *TeamGrade) GetOffenseGradeLetter() string {
	return tg.OffenseGradeLetter
}

func (tg *TeamGrade) GetDefenseGradeLetter() string {
	return tg.DefenseGradeLetter
}

func (tg *TeamGrade) GetSpecialTeamsGradeLetter() string {
	return tg.SpecialTeamsGradeLetter
}

func (tg *TeamGrade) GetOverallGradeLetter() string {
	return tg.OverallGradeLetter
}
