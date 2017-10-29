/*
* @Author: huang
* @Date:   2017-10-25 17:23:41
* @Last Modified by:   huang
* @Last Modified time: 2017-10-25 17:25:59
 */
package dbmgr

import (
	"comm/db"
	"fmt"
	"time"
)

// ============================================================================

type seqid_t struct {
	Id        int   `bson:"_id"`
	ArticleId int64 `bson:"articleid"`
}

// ============================================================================

func CenterCreateSeqId() {
	if DBCenter.HasCollection(CTableSeqId) {
		return
	}

	var obj seqid_t

	obj.Id = 1
	obj.ArticleId = 999999

	err := DBCenter.Insert(CTableSeqId, &obj)
	if err != nil {
		log.Error("dbmgr.Center_CreateSeqId() failed:", err)
	}
}

func GenArticleId() string {
	var obj seqid_t

	err := DBCenter.FindAndModify(
		CTableSeqId,
		db.M{"_id": 1},
		db.Change{
			Update: db.M{
				"$inc": db.M{"articleid": 1},
			},
			ReturnNew: true,
		},
		db.M{"articleid": 1},
		&obj,
	)
	if err != nil {
		log.Error("dbmgr.Center_GenUserId() failed:", err)
	}

	return fmt.Sprintf("%d%d", time.Now().Unix(), obj.ArticleId)
}
