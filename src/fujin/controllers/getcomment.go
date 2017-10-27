/*
* @Author: huang
* @Date:   2017-10-26 16:25:55
* @Last Modified by:   huang
* @Last Modified time: 2017-10-27 16:03:26
 */
package controllers

import (
	"comm/dbmgr"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ============================================================================

type GetCommentReq struct {
	SessionKey string    `json:"sessionkey"` //session_key
	Uid        string    `json:"uid"`        //自己的uid
	ArticleId  string    `json:"articleid"`  //评论的文章Id
	Loc        *Location `json:"loc"`        //写的位置
}

type CommentOne struct {
	CName    string `json:"name"`     //评论者名字
	CHead    string `json:"head"`     //评论者头像
	Distance int32  `json:"distance"` //写的距离
	Ts       int64  `json:"ts"`       //写时间
	Content  string `bson:"content"`  //内容
}

type GetCommentRes struct {
	N   int32         `json:n`     //总评论数
	Cmt []*CommentOne `json:"cmt"` //评论组
}

// ============================================================================

func GetCommentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	if r.Method != "POST" {
		return
	}

	result, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	var req GetCommentReq
	err := json.Unmarshal([]byte(result), &req)
	log.Info(req)

	if err != nil {
		log.Error("json Unmarshal error", result, err)
		w.Write([]byte(ErrGetCommentFailed))
		return
	}

	if len(req.Loc.Coordinates) != 2 {
		log.Error("Coordinates error", req.Loc.Coordinates)
		w.Write([]byte(ErrGetCommentFailed))
		return
	}

	if !CheckSessionKey(req.Uid, req.SessionKey) {
		log.Error("CheckSessionKey error", req.Uid, req.SessionKey)
		w.Write([]byte(ErrGetCommentFailed))
		return
	}

	cmts := dbmgr.GetComments(req.ArticleId)
	if cmts == nil {
		w.Write([]byte(ErrGetCommentFailed))
		return
	}

	res := &GetCommentRes{}
	isRes := false
	for _, v := range cmts.Cmt {
		log.Info(v)

		one := &CommentOne{}
		one.CName = v.CName
		one.CHead = v.CHead
		one.Content = v.Content
		one.Distance = int32(EarthDistance(req.Loc.Coordinates[0], req.Loc.Coordinates[1], v.Loc.Coordinates[0], v.Loc.Coordinates[1]))
		one.Ts = v.Ts.Unix()

		res.Cmt = append(res.Cmt, one)
		isRes = true
	}

	res.N = cmts.CmtCnt

	if !isRes {
		w.Write([]byte(Success))
		return
	}

	b, err := json.Marshal(res)
	if err != nil {
		log.Error("json.Marshal(res) error")
		w.Write([]byte(ErrGetCommentFailed))
		return
	}

	w.Write([]byte(string(b)))
}
