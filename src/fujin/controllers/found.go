/*
* @Author: huang
* @Date:   2017-10-26 14:14:30
* @Last Modified by:   huang
* @Last Modified time: 2018-05-08 11:36:48
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

type FoundReq struct {
	SessionKey string    `json:"sessionkey"` //session_key
	Uid        string    `json:"uid"`
	IsSelf     bool      `json:"isself"`
	Loc        *Location `json:"loc"`      //写的位置
	ReqIndex   int32     `json:"reqindex"` //请求数量index
}

type articleOneRes struct {
	Id          string   `json:"id"`         //文章id
	AuthorName  string   `json:"authorname"` //作者名字
	AuthorHead  string   `json:"authorhead"` //作者头像
	Content     string   `json:"content"`    //内容
	Images      []string `json:"images"`     //图像地址
	Distance    string   `json:"distance"`   //距离
	Ts          string   `json:"ts"`         //时间
	CommentsNum int32    `json:"cmtnum"`     //评论条数
}

type FoundRes struct {
	ErrorCode int32            `json:"errorcode"` //错误码
	Articles  []*articleOneRes `json:"articles"`
}

// ============================================================================

func FoundHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	if r.Method != "POST" {
		return
	}

	result, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	var req FoundReq
	err := json.Unmarshal([]byte(result), &req)
	log.Info(req)

	if err != nil {
		log.Error("json Unmarshal error", result, err)
		w.Write([]byte(ErrFoundFailed))
		return
	}

	if len(req.Loc.Coordinates) != 2 {
		log.Error("Coordinates error", req.Loc.Coordinates)
		w.Write([]byte(ErrFoundFailed))
		return
	}

	if !CheckSessionKey(req.Uid, req.SessionKey) {
		log.Error("CheckSessionKey error", req.Uid, req.SessionKey)
		w.Write([]byte(ErrTokenError))
		return
	}

	var arr []*dbmgr.Articles
	if req.IsSelf {
		//arr = dbmgr.GetArticlesByAuthorId(req.Uid)
		arr = dbmgr.GetArticlesByAuthorIdLimit(req.Uid, int(req.ReqIndex), int(config.Common.PerReqNum))

	} else {
		//arr = dbmgr.GetArticlesByLocation(req.Loc.Coordinates[0], req.Loc.Coordinates[1], 0)
		arr = dbmgr.GetArticlesByLocationByLimit(req.Loc.Coordinates[0], req.Loc.Coordinates[1], 0, int(req.ReqIndex), int(config.Common.PerReqNum))
	}

	isRes := false
	res := &FoundRes{}
	for _, v := range arr {
		log.Info(v)

		one := &articleOneRes{}
		one.Id = v.Id
		one.AuthorName = v.AuthorName
		one.AuthorHead = v.AuthorHead
		one.Content = wordsfilter.Filter(v.Content)
		one.Images = v.Images
		one.Distance = EarthDistance(req.Loc.Coordinates[0], req.Loc.Coordinates[1], v.Loc.Coordinates[0], v.Loc.Coordinates[1])
		one.Ts = TimeGapStr(v.Ts)
		one.CommentsNum = dbmgr.GetCommentsNum(v.Id)

		res.Articles = append(res.Articles, one)
		isRes = true
	}
	res.ErrorCode = 200

	if !isRes {
		w.Write([]byte(Success))
		return
	}

	b, err := json.Marshal(res)
	if err != nil {
		log.Error("json.Marshal(res) error")
		w.Write([]byte(ErrFoundFailed))
		return
	}

	w.Write([]byte(string(b)))
}
