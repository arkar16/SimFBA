package repository

import (
	"log"

	"github.com/CalebRose/SimFBA/models"
	"github.com/CalebRose/SimFBA/structs"
	"gorm.io/gorm"
)

func CreateCFBPlayByPlaysInBatch(plays []structs.CollegePlayByPlay, db *gorm.DB) {
	err := db.CreateInBatches(&plays, len(plays)).Error
	if err != nil {
		log.Panicln("Could not save play by plays!")
	}
}

func CreateNFLPlayByPlaysInBatch(plays []structs.NFLPlayByPlay, db *gorm.DB) {
	err := db.CreateInBatches(&plays, len(plays)).Error
	if err != nil {
		log.Panicln("Could not save play by plays!")
	}
}

func CreateCFBTeamStatsInBatch(stats []structs.CollegeTeamStats, db *gorm.DB) {
	err := db.CreateInBatches(&stats, len(stats)).Error
	if err != nil {
		log.Panicln("Could not save team stats!")
	}
}

func CreateCFBPlayerStatsInBatch(stats []structs.CollegePlayerStats, db *gorm.DB) {
	err := db.CreateInBatches(&stats, len(stats)).Error
	if err != nil {
		log.Panicln("Could not save team stats!")
	}
}

func CreateNFLTeamStatsInBatch(stats []structs.NFLTeamStats, db *gorm.DB) {
	err := db.CreateInBatches(&stats, len(stats)).Error
	if err != nil {
		log.Panicln("Could not save team stats!")
	}
}

func CreateNFLPlayerStatsInBatch(stats []structs.NFLPlayerStats, db *gorm.DB) {
	err := db.CreateInBatches(&stats, len(stats)).Error
	if err != nil {
		log.Panicln("Could not save team stats!")
	}
}

func CreateRecruitProfileRecord(cp structs.RecruitPlayerProfile, db *gorm.DB) {
	err := db.Create(&cp).Error
	if err != nil {
		log.Panicln("Could not create recruit profile record!")
	}
}

func CreateCFBSnapsInBatch(snaps []structs.CollegePlayerGameSnaps, db *gorm.DB) {
	err := db.CreateInBatches(&snaps, len(snaps)).Error
	if err != nil {
		log.Panicln("Could not create college snaps in batch!")
	}
}

func CreateNFLSnapsInBatch(snaps []structs.NFLPlayerGameSnaps, db *gorm.DB) {
	err := db.CreateInBatches(&snaps, len(snaps)).Error
	if err != nil {
		log.Panicln("Could not create nfl snaps in batch!")
	}
}

func CreateCFBSeasonSnaps(snap structs.CollegePlayerSeasonSnaps, db *gorm.DB) {
	err := db.Create(&snap).Error
	if err != nil {
		log.Panicln("Could not create cfb season snaps record!")
	}
}

func CreateNFLSeasonSnaps(snap structs.NFLPlayerSeasonSnaps, db *gorm.DB) {
	err := db.Create(&snap).Error
	if err != nil {
		log.Panicln("Could not create nfl season snaps record!")
	}
}

func CreateCollegePromiseRecord(promise structs.CollegePromise, db *gorm.DB) {
	// Save College Player Record
	err := db.Create(&promise).Error
	if err != nil {
		log.Panicln("Could not save new college recruit record")
	}
}

func CreateNotification(noti structs.Notification, db *gorm.DB) {
	err := db.Create(&noti).Error
	if err != nil {
		log.Panicln("Could not create notification record!")
	}
}

func CreateCFBPlayerRecord(player structs.CollegePlayer, db *gorm.DB) {
	err := db.Create(&player).Error
	if err != nil {
		log.Panicln("Could not create cfb season snaps record!")
	}
}

func CreateHistoricCFBPlayerRecord(player structs.HistoricCollegePlayer, db *gorm.DB) {
	err := db.Create(&player).Error
	if err != nil {
		log.Panicln("Could not create cfb season snaps record!")
	}
}

func CreateNFLDrafteeRecord(player models.NFLDraftee, db *gorm.DB) {
	err := db.Create(&player).Error
	if err != nil {
		log.Panicln("Could not create cfb season snaps record!")
	}
}

func CreateRetireeRecord(player structs.NFLRetiredPlayer, db *gorm.DB) {
	err := db.Create(&player).Error
	if err != nil {
		log.Panicln("Could not create cfb season snaps record!")
	}
}

func CreateNFLDrafteesInBatches(db *gorm.DB, draftees []models.NFLDraftee, batchSize int) error {
	// Create the records in batches with the specified batch size
	if err := db.CreateInBatches(draftees, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func CreateCFBRecruitRecordsBatch(db *gorm.DB, croots []structs.Recruit, batchSize int) error {
	total := len(croots)
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		if err := db.CreateInBatches(croots[i:end], batchSize).Error; err != nil {
			return err
		}
	}
	return nil
}

func CreateGlobalPlayerRecordsBatch(db *gorm.DB, croots []structs.Player, batchSize int) error {
	total := len(croots)
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		if err := db.CreateInBatches(croots[i:end], batchSize).Error; err != nil {
			return err
		}
	}
	return nil
}

func CreateNFLDrafteesSafely(db *gorm.DB, draftees []models.NFLDraftee, batchSize int) error {
	total := len(draftees)
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		if err := db.CreateInBatches(draftees[i:end], batchSize).Error; err != nil {
			return err
		}
	}
	return nil
}

func CreateTransferPortalProfileRecord(profile structs.TransferPortalProfile, db *gorm.DB) {
	err := db.Create(&profile).Error
	if err != nil {
		log.Panicln("Could not create cfb season snaps record!")
	}
}

func CreateTransferPortalProfileRecordsBatch(db *gorm.DB, profiles []structs.TransferPortalProfile, batchSize int) error {
	total := len(profiles)
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		if err := db.CreateInBatches(profiles[i:end], batchSize).Error; err != nil {
			return err
		}
	}
	return nil
}
