/*
* @Author: huang
* @Date:   2017-10-26 16:25:55
* @Last Modified by:   huang
* @Last Modified time: 2017-11-01 10:32:53
 */
package controllers

import (
	"comm/config"
	"comm/dbmgr"
	"comm/wordsfilter"
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
	ReqIndex   int32     `json:"reqindex"`   //请求数量index
}

type CommentOne struct {
	Id       string `json:"id"`       //评论id
	CName    string `json:"name"`     //评论者名字
	CHead    string `json:"head"`     //评论者头像
	Distance string `json:"distance"` //写的距离
	Ts       string `json:"ts"`       //写时间
	Content  string `json:"content"`  //内容
}

type GetCommentRes struct {
	ErrorCode int32         `json:"errorcode"` //错误码
	N         int32         `json:n`           //总评论数
	Cmt       []*CommentOne `json:"cmt"`       //评论组
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

	//cmts := dbmgr.GetComments(req.ArticleId)
	cmts := dbmgr.GetCommentsByLimit(req.ArticleId, int(req.ReqIndex), int(config.Common.PerReqNum))
	if cmts == nil {
		w.Write([]byte(ErrGetCommentFailed))
		return
	}

	res := &GetCommentRes{}
	isRes := false
	for _, v := range cmts.Cmt {
		log.Info(v)

		one := &CommentOne{}
		one.Id = v.Id
		one.CName = v.CName
		one.CHead = v.CHead
		one.Content = wordsfilter.Filter(v.Content)
		one.Distance = EarthDistance(req.Loc.Coordinates[0], req.Loc.Coordinates[1], v.Loc.Coordinates[0], v.Loc.Coordinates[1])
		one.Ts = TimeGapStr(v.Ts)

		res.Cmt = append(res.Cmt, one)
		isRes = true
	}

	res.N = cmts.CmtCnt
	res.ErrorCode = 200

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
