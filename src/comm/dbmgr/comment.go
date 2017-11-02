/*
* @Author: huang
* @Date:   2017-10-26 15:26:00
* @Last Modified by:   huang
* @Last Modified time: 2017-11-01 17:34:41
 */
package dbmgr

import (
	"comm/db"
	"time"
)

// ============================================================================
type CommentOne struct {
	Id        string    `bson:"id"`      //评论id
	CUid      string    `bson:"cuid"`    //评论者Id
	CName     string    `bson:"cname"`   //评论者名字
	CHead     string    `bson:"chead"`   //评论者头像
	Loc       *Location `bson:"loc"`     //写的位置
	Ts        time.Time `bson:"ts"`      //写时间
	Content   string    `bson:"content"` //内容
	Anonymous bool      `bson:"anon"`    //是否匿名
}

type Comments struct {
	Id     string        `bson:"_id"`    //文章Id
	CmtCnt int32         `bson:"cmtcnt"` //评论数量
	Cmt    []*CommentOne `bson:"cmt"`    //评论组
}

// ============================================================================

func GetComments(id string) *Comments {
	var obj Comments

	err := DBCenter.GetObjectByCond(
		CTableComments,
		db.M{
			"_id": id,
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

func GetCommentsNum(id string) int32 {
	var obj Comments

	err := DBCenter.GetObjectByCond(
		CTableComments,
		db.M{
			"_id": id,
		},
		&obj,
	)

	if err == nil {
		return obj.CmtCnt
	} else {
		// failed
		return 0
	}
}

func GetCommentsByLimit(id string, skip, limit int) *Comments {
	arr := []int{skip, limit}

	var obj Comments

	err := DBCenter.GetProjectionByCond(
		CTableComments,
		db.M{
			"_id": id,
		},
		db.M{
			"cmt": db.M{
				"$slice": arr,
			},
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

func WriteComment(id string, cmt *CommentOne) {
	if cmt == nil {
		return
	}

	var obj CommentOne

	obj.Id = GenCommentId()
	obj.CUid = cmt.CUid
	obj.CName = cmt.CName
	obj.CHead = cmt.CHead
	obj.Loc = cmt.Loc
	obj.Ts = cmt.Ts
	obj.Content = cmt.Content
	obj.Anonymous = cmt.Anonymous

	err := DBCenter.Upsert(
		CTableComments,
		id,
		db.M{
			"$push": db.M{
				"cmt": &obj,
			},
			"$inc": db.M{"cmtcnt": 1},
		},
	)

	if err != nil {
		log.Error(err)
	}
}
