package repository

import (
	"github.com/CalebRose/SimFBA/dbprovider"
	"github.com/CalebRose/SimFBA/structs"
)

func FindRecruitPlayerProfileRecords(profileID, recruitID string, includeRecruit, orderByPoints, removeFromBoard bool) []structs.RecruitPlayerProfile {
	db := dbprovider.GetInstance().GetDB()

	var croots []structs.RecruitPlayerProfile

	query := db.Model(&croots)

	if includeRecruit {
		query = query.Preload("Recruit")
	}

	if len(profileID) > 0 {
		query = query.Where("profile_id = ?", profileID)
	}

	if len(recruitID) > 0 {
		query = query.Where("recruit_id = ?", recruitID)
	}

	if removeFromBoard {
		query = query.Where("removed_from_board = ?", false)
	}

	if orderByPoints {
		query = query.Order("total_points desc")
	}

	if err := query.Find(&croots).Error; err != nil {
		return []structs.RecruitPlayerProfile{}
	}

	return croots
}
