/*
* @Author: huang
* @Date:   2017-10-26 15:26:00
* @Last Modified by:   huang
* @Last Modified time: 2018-05-09 14:52:39
 */
package dbmgr

import (
	"comm/db"
	"time"
)

// ============================================================================
// 点赞
type ThumbOne struct {
	Uid string `bson:"uid"`
}

type CommentOne struct {
	Id        string      `bson:"id"`      //评论id
	CUid      string      `bson:"cuid"`    //评论者Id
	CName     string      `bson:"cname"`   //评论者名字
	CHead     string      `bson:"chead"`   //评论者头像
	Loc       *Location   `bson:"loc"`     //写的位置
	Ts        time.Time   `bson:"ts"`      //写时间
	Content   string      `bson:"content"` //内容
	Thumb     []*ThumbOne `bson:"thumb"`   //点赞信息
	Anonymous bool        `bson:"anon"`    //是否匿名
}

type Comments struct {
	Id     string        `bson:"_id"`    //文章Id
	CmtCnt int32         `bson:"cmtcnt"` //评论数量
	Cmt    []*CommentOne `bson:"cmt"`    //评论组
	Thumb  []*ThumbOne   `bson:"thumb"`  //点赞信息
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

// 更新评论条数
func UpdateCommentNum(id string) {
	var obj Comments

	err := DBCenter.GetObjectByCond(
		CTableComments,
		db.M{
			"_id": id,
		},
		&obj,
	)

	cnt := 0
	if err == nil {
		cnt = len(obj.Cmt)
	} else {
		return
	}

	DBCenter.UpdateByCond(
		CTableComments,
		db.M{
			"_id": id,
		},
		db.M{
			"$set": db.M{"cmtcnt": cnt},
		},
	)
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

// 写评论
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

// 删除评论
func CenterDelComment(authorid string, articleid string, commentid string) error {
	err := DBCenter.UpdateByCond(
		CTableComments,
		db.M{
			"_id": articleid,
		},
		db.M{
			"$pull": db.M{
				"cmt": db.M{
					"id":   commentid,
					"cuid": authorid,
				},
			},
		},
	)

	if err != nil {
		log.Warning("CenterDelComment failed:", err)
	}

	// 相对比较低频的操作
	UpdateCommentNum(articleid)

	return err
}

// 文章点赞
func ArticleThumbAdd(uid string, articleid string) error {
	th := &ThumbOne{Uid: uid}

	err := DBCenter.UpdateByCond(
		CTableComments,
		db.M{
			"_id": articleid,
		},
		db.M{
			"$addToSet": db.M{
				"thumb": th,
			},
		},
	)

	if err != nil {
		log.Warning("ArticleThumbAdd failed:", err)
	}

	return err
}

// 文章点赞数，自己是否点赞
func ArticleThumbNum(uid string, articleid string) (int32, bool) {
	var obj Comments

	err := DBCenter.GetObjectByCond(
		CTableComments,
		db.M{
			"_id": articleid,
		},
		&obj,
	)

	if err != nil {
		log.Warning("ArticleThumbNum failed:", err)
		return 0, false
	}

	for _, v := range obj.Thumb {
		if v.Uid == uid {
			return int32(len(obj.Thumb)), true
		}
	}

	return int32(len(obj.Thumb)), false
}

// 评论点赞
func CommentThumbAdd(uid string, articleid string, commentid string) error {
	th := &ThumbOne{Uid: uid}

	err := DBCenter.UpdateByCond(
		CTableComments,
		db.M{
			"_id":    articleid,
			"cmt.id": commentid,
		},
		db.M{
			"$addToSet": db.M{
				"cmt.$.thumb": th,
			},
		},
	)

	if err != nil {
		log.Warning("CommentThumbAdd failed:", err)
	}

	return err
}

// 评论点赞数，自己是否点赞
func CommentThumbNum(uid string, articleid string, commentid string) (int32, bool) {
	var obj Comments

	err := DBCenter.GetObjectByCond(
		CTableComments,
		db.M{
			"_id":    articleid,
			"cmt.id": commentid,
		},
		&obj,
	)

	if err != nil {
		log.Warning("CommentThumbNum failed:", err)
		return 0, false
	}

	for _, v := range obj.Cmt {
		if v.Id == commentid {
			for _, c := range v.Thumb {
				if c.Uid == uid {
					return int32(len(v.Thumb)), true
				}
			}

			return int32(len(v.Thumb)), false
		}
	}

	return 0, false
}
