package structs

type TeamGrade struct {
	OffenseGradeNumber      float32
	DefenseGradeNumber      float32
	SpecialTeamsGradeNumber float32
	OverallGradeNumber      float32
	OffenseGradeLetter      string
	DefenseGradeLetter      string
	SpecialTeamsGradeLetter string
	OverallGradeLetter      string
}

func (tg *TeamGrade) SetOffenseGradeNumber(grade float32) {
	tg.OffenseGradeNumber = grade
}

func (tg *TeamGrade) SetDefenseGradeNumber(grade float32) {
	tg.DefenseGradeNumber = grade
}

func (tg *TeamGrade) SetSpecialTeamsGradeNumber(grade float32) {
	tg.SpecialTeamsGradeNumber = grade
}

func (tg *TeamGrade) SetOverallGradeNumber(grade float32) {
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

func (tg *TeamGrade) GetOffenseGradeNumber() float32 {
	return tg.OffenseGradeNumber
}

func (tg *TeamGrade) GetDefenseGradeNumber() float32 {
	return tg.DefenseGradeNumber
}

func (tg *TeamGrade) GetSpecialTeamsGradeNumber() float32 {
	return tg.SpecialTeamsGradeNumber
}

func (tg *TeamGrade) GetOverallGradeNumber() float32 {
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
