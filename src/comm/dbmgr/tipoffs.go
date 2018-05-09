/*
* @Author: huang
* @Date:   2018-05-09 16:31:15
* @Last Modified by:   huang
* @Last Modified time: 2018-05-09 16:41:29
 */
package dbmgr

// ============================================================================

type TipOffs struct {
	Uid       string `bson:"uid"`
	Name      string `bson:"name"`
	Head      string `bson:"head"`
	TipType   string `bson:"tiptype"`   //举报类型
	ArticleId string `bson:"articleid"` //举报文章id
	CommentId string `bson:"commentid"` //举报评论id
	Content   string `bson:"content"`   //内容
}

// ============================================================================

func InsertTipOffs(uid, name, head, tiptype, articleid, commentid, content string) {
	obj := TipOffs{
		Uid:       uid,
		Name:      name,
		Head:      head,
		TipType:   tiptype,
		ArticleId: articleid,
		CommentId: commentid,
		Content:   content,
	}

	// save to db
	err := DBCenter.Insert(CTableTipOffs, &obj)
	if err != nil {
		log.Error("create user failed:", uid, err)
	}
}
