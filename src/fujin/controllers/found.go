/*
* @Author: huang
* @Date:   2017-10-26 14:14:30
* @Last Modified by:   huang
* @Last Modified time: 2017-10-27 16:03:59
 */
package controllers

import (
	"comm/dbmgr"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ============================================================================

type FoundReq struct {
	SessionKey string    `json:"sessionkey"` //session_key
	Uid        string    `json:"uid"`
	IsSelf     bool      `json:"isself"`
	Loc        *Location `json:"loc"` //写的位置
}

type articleOneRes struct {
	Id         string   `json:id`           //文章id
	AuthorName string   `json:"authorname"` //作者名字
	AuthorHead string   `json:"authorhead"` //作者头像
	Content    string   `json:"content"`    //内容
	Images     []string `json:"images"`     //图像地址
	Distance   int32    `json:"distance"`   //距离
	Ts         string   `json:"ts"`         //时间
}

type FoundRes struct {
	Articles []*articleOneRes `json:"articles"`
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
		w.Write([]byte(ErrFoundFailed))
		return
	}

	var arr []*dbmgr.Articles
	if req.IsSelf {
		arr = dbmgr.GetArticlesByAuthorId(req.Uid)
	} else {
		arr = dbmgr.GetArticlesByLocation(req.Loc.Coordinates[0], req.Loc.Coordinates[1], 0)
	}

	isRes := false
	res := &FoundRes{}
	for _, v := range arr {
		log.Info(v)

		one := &articleOneRes{}
		one.Id = v.Id
		one.AuthorName = v.AuthorName
		one.AuthorHead = v.AuthorHead
		one.Content = v.Content
		one.Images = v.Images
		one.Distance = int32(EarthDistance(req.Loc.Coordinates[0], req.Loc.Coordinates[1], v.Loc.Coordinates[0], v.Loc.Coordinates[1]))
		one.Ts = TimeGapStr(v.Ts)

		res.Articles = append(res.Articles, one)
		isRes = true
	}

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
