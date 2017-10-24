package dbmgr

import (
	"comm/db"
	"time"
)

// ============================================================================

type Users struct {
	Uid        string    `bson:"_id"`        //用户唯一id
	SessionKey string    `bson:"sessionkey"` //回话key
	Expire     time.Time `bson:"expire"`     //回话超时时间
}

// ============================================================================

func CenterGetUserInfo(uid string) *Users {
	var obj Users

	err := DBCenter.GetObjectByCond(
		CTableUsers,
		db.M{
			"_id": uid,
		},
		&obj,
	)

	if err == nil {

		return &obj

	} else {
		// failed
		return nil
	}
}

func CenterUpdateUserInfo(uid string, sessionkey string, expire time.Time) {

	err := DBCenter.Upsert(
		CTableUsers,
		uid,
		db.M{
			"$set": db.M{
				"sessionkey": sessionkey,
				"expire":     expire,
			},
		},
	)

	if err != nil {
		log.Warning("save arena rank data failed:", err)
	}

	return
}

// ============================================================================
