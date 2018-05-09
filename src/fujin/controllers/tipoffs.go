/*
* @Author: huang
* @Date:   2018-05-09 16:38:47
* @Last Modified by:   huang
* @Last Modified time: 2018-05-09 16:42:38
 */

package controllers

import (
	"comm/dbmgr"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ============================================================================

type TipOffsReq struct {
	Uid       string `bson:"uid"`
	Name      string `bson:"name"`
	Head      string `bson:"head"`
	TipType   string `bson:"tiptype"`   //举报类型
	ArticleId string `bson:"articleid"` //举报文章id
	CommentId string `bson:"commentid"` //举报评论id
	Content   string `bson:"content"`   //内容
}

// ============================================================================

func TipOffsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	if r.Method != "POST" {
		return
	}

	result, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	var req TipOffsReq
	err := json.Unmarshal([]byte(result), &req)
	log.Info(req)

	if err != nil {
		log.Error("json Unmarshal error", result, err)
		w.Write([]byte(ErrCallBackFailed))
		return
	}

	dbmgr.InsertTipOffs(req.Uid, req.Name, req.Head, req.TipType, req.ArticleId, req.CommentId, req.Content)

	w.Write([]byte(Success))
}
