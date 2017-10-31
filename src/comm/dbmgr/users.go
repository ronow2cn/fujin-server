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
	UnionId    string    `bson:"unionid"`    //unionid
	Name       string    `bson:"name"`       //名字
	Head       string    `bson:"head"`       //头像
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

func CenterUpdateUserInfo(uid string, sessionkey string, expire time.Time, unionid string) error {
	obj := &Users{
		Uid:        uid,
		SessionKey: sessionkey,
		Expire:     expire,
		UnionId:    unionid,
	}

	err := DBCenter.Upsert(
		CTableUsers,
		uid,
		obj,
	)

	if err != nil {
		log.Warning("save user data failed:", err)
	}

	return err
}

func CenterUpdateUserNameHead(uid string, name, head string) error {
	err := DBCenter.Update(
		CTableUsers,
		uid,
		db.M{
			"$set": db.M{
				"name": name,
				"head": head,
			},
		},
	)

	if err != nil {
		log.Warning("update user name head failed:", err)
	}

	return err
}

// ============================================================================
