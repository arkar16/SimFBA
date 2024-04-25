package structs

type InjuryReportResponse struct {
	NFLPlayers     []NFLPlayer
	CollegePlayers []CollegePlayer
}

type SimCFBStatsResponse struct {
	CollegeConferences []CollegeConference
	CollegePlayers     []CollegePlayerResponse
	CollegeTeams       []CollegeTeamResponse
}

type SimNFLStatsResponse struct {
	NFLPlayers []NFLPlayerResponse
	NFLTeams   []NFLTeamResponse
}

type CollegeTeamResponse struct {
	ID int
	BaseTeam
	ConferenceID int
	Conference   string
	DivisionID   int
	Division     string
	Stats        CollegeTeamStats
	SeasonStats  CollegeTeamSeasonStats
}

type NFLTeamResponse struct {
	ID int
	BaseTeam
	ConferenceID int
	Conference   string
	DivisionID   int
	Division     string
	Stats        NFLTeamStats
	SeasonStats  NFLTeamSeasonStats
}

type CrootProfile struct {
	ID                        uint
	SeasonID                  int
	RecruitID                 int
	ProfileID                 int
	TotalPoints               float64
	CurrentWeeksPoints        float64
	SpendingCount             int
	RecruitingEfficiencyScore float64
	Scholarship               bool
	ScholarshipRevoked        bool
	AffinityOneEligible       bool
	AffinityTwoEligible       bool
	TeamAbbreviation          string
	RemovedFromBoard          bool
	IsSigned                  bool
	IsLocked                  bool
	CaughtCheating            bool
	Recruit                   Croot
}

func (cp *CrootProfile) Map(rp RecruitPlayerProfile, c Croot) {
	cp.ID = rp.ID
	cp.SeasonID = rp.SeasonID
	cp.RecruitID = rp.RecruitID
	cp.ProfileID = rp.ProfileID
	cp.TotalPoints = rp.TotalPoints
	cp.CurrentWeeksPoints = rp.CurrentWeeksPoints
	cp.SpendingCount = rp.SpendingCount
	cp.RecruitingEfficiencyScore = rp.RecruitingEfficiencyScore
	cp.Scholarship = rp.Scholarship
	cp.ScholarshipRevoked = rp.ScholarshipRevoked
	cp.AffinityOneEligible = rp.AffinityOneEligible
	cp.AffinityTwoEligible = rp.AffinityTwoEligible
	cp.TeamAbbreviation = rp.TeamAbbreviation
	cp.RemovedFromBoard = rp.RemovedFromBoard
	cp.IsSigned = rp.IsSigned
	cp.IsLocked = rp.IsLocked
	cp.CaughtCheating = rp.CaughtCheating
	cp.Recruit = c
}

// Sorting Funcs
type ByCrootProfileTotal []CrootProfile

func (rp ByCrootProfileTotal) Len() int      { return len(rp) }
func (rp ByCrootProfileTotal) Swap(i, j int) { rp[i], rp[j] = rp[j], rp[i] }
func (rp ByCrootProfileTotal) Less(i, j int) bool {
	return rp[i].TotalPoints > rp[j].TotalPoints
}

type SimTeamBoardResponse struct {
	ID                        uint
	TeamID                    int
	Team                      string
	TeamAbbreviation          string
	State                     string
	ScholarshipsAvailable     int
	WeeklyPoints              float64
	SpentPoints               float64
	TotalCommitments          int
	RecruitClassSize          int
	BaseEfficiencyScore       float64
	RecruitingEfficiencyScore float64
	PreviousOverallWinPer     float64
	PreviousConferenceWinPer  float64
	CurrentOverallWinPer      float64
	CurrentConferenceWinPer   float64
	ESPNScore                 float64
	RivalsScore               float64
	Rank247Score              float64
	CompositeScore            float64
	IsAI                      bool
	IsUserTeam                bool
	BattlesWon                int
	BattlesLost               int
	AIMinThreshold            int
	AIMaxThreshold            int
	AIStarMin                 int
	AIStarMax                 int
	AIAutoOfferscholarships   bool
	OffensiveScheme           string
	DefensiveScheme           string
	Recruiter                 string
	RecruitingClassRank       int
	Recruits                  []CrootProfile
	Affinities                []ProfileAffinity `gorm:"foreignKey:ProfileID"`
}

func (stbr *SimTeamBoardResponse) Map(rtp RecruitingTeamProfile, c []CrootProfile) {
	stbr.ID = rtp.ID
	stbr.TeamID = rtp.TeamID
	stbr.Team = rtp.Team
	stbr.IsAI = rtp.IsAI
	stbr.TeamAbbreviation = rtp.TeamAbbreviation
	stbr.State = rtp.State
	stbr.ScholarshipsAvailable = rtp.ScholarshipsAvailable
	stbr.WeeklyPoints = rtp.WeeklyPoints
	stbr.SpentPoints = rtp.SpentPoints
	stbr.TotalCommitments = rtp.TotalCommitments
	stbr.BaseEfficiencyScore = rtp.BaseEfficiencyScore
	stbr.RecruitingEfficiencyScore = rtp.RecruitingEfficiencyScore
	stbr.ESPNScore = rtp.ESPNScore
	stbr.RivalsScore = rtp.RivalsScore
	stbr.Rank247Score = rtp.Rank247Score
	stbr.CompositeScore = rtp.CompositeScore
	stbr.RecruitingClassRank = rtp.RecruitingClassRank
	stbr.Affinities = rtp.Affinities
	stbr.Recruits = c
	stbr.RecruitClassSize = rtp.RecruitClassSize
	stbr.IsUserTeam = rtp.IsUserTeam
	stbr.BattlesWon = rtp.BattlesWon
	stbr.BattlesLost = rtp.BattlesLost
	stbr.AIMinThreshold = rtp.AIMinThreshold
	stbr.AIMaxThreshold = rtp.AIMaxThreshold
	stbr.AIStarMin = rtp.AIStarMin
	stbr.AIStarMax = rtp.AIStarMax
	stbr.OffensiveScheme = rtp.OffensiveScheme
	stbr.DefensiveScheme = rtp.DefensiveScheme
	stbr.Recruiter = rtp.Recruiter
}

type DashboardTeamProfileResponse struct {
	TeamProfile  RecruitingTeamProfile
	TeamNeedsMap map[string]int
}

func (d *DashboardTeamProfileResponse) SetTeamProfile(profile RecruitingTeamProfile) {
	d.TeamProfile = profile
}

func (d *DashboardTeamProfileResponse) SetTeamNeedsMap(obj map[string]int) {
	d.TeamNeedsMap = obj
}

type TeamBoardTeamProfileResponse struct {
	TeamProfile  SimTeamBoardResponse
	TeamNeedsMap map[string]int
}

func (t *TeamBoardTeamProfileResponse) SetTeamProfile(profile SimTeamBoardResponse) {
	t.TeamProfile = profile
}

func (t *TeamBoardTeamProfileResponse) SetTeamNeedsMap(obj map[string]int) {
	t.TeamNeedsMap = obj
}
