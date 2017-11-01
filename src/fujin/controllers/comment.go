/*
* @Author: huang
* @Date:   2017-10-26 15:22:38
* @Last Modified by:   huang
* @Last Modified time: 2017-11-01 09:59:03
 */
package controllers

import (
	"comm/dbmgr"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// ============================================================================

type CommentReq struct {
	SessionKey string    `json:"sessionkey"` //session_key
	Uid        string    `json:"uid"`        //自己的uid
	Name       string    `json:"name"`       //名字
	Head       string    `json:"head"`       //头像
	Loc        *Location `json:"loc"`        //写的位置
	ArticleId  string    `json:"articleid"`  //要评论的文章Id
	Content    string    `json:"content"`    //内容
	Anonymous  bool      `json:"anon"`       //是否匿名
}

// ============================================================================

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	if r.Method != "POST" {
		return
	}

	result, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	var req CommentReq
	err := json.Unmarshal([]byte(result), &req)
	log.Info(req)

	if err != nil {
		log.Error("json Unmarshal error", result, err)
		w.Write([]byte(ErrCommentFailed))
		return
	}

	if !CheckSessionKey(req.Uid, req.SessionKey) {
		log.Error("CheckSessionKey error", req.Uid, req.SessionKey)
		w.Write([]byte(ErrCommentFailed))
		return
	}

	if len(req.Loc.Coordinates) != 2 {
		log.Error("Coordinates error", req.Loc.Coordinates)
		w.Write([]byte(ErrCommentFailed))
		return
	}

	dbmgr.WriteComment(req.ArticleId, &dbmgr.CommentOne{
		CUid:  req.Uid,
		CName: req.Name,
		CHead: req.Head,
		Loc: &dbmgr.Location{
			Type:        "Point",
			Coordinates: req.Loc.Coordinates,
		},
		Ts:        time.Now(),
		Content:   req.Content,
		Anonymous: req.Anonymous,
	})

	w.Write([]byte(Success))

	//只保存发言的用户名和头像
	err = dbmgr.CenterUpdateUserNameHead(req.Uid, req.Name, req.Head)
	if err != nil {
		log.Error("CenterUpdateUserInfo error", err)
		return
	}
}
