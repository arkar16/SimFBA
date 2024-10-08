package controller

import (
	"github.com/CalebRose/SimFBA/managers"
	"github.com/CalebRose/SimFBA/structs"
)

func GetUpdatedTimestamp() structs.Timestamp {
	ts := managers.GetTimestamp()
	return ts
}
