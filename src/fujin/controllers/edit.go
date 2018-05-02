/*
* @Author: huang
* @Date:   2017-10-26 10:03:30
* @Last Modified by:   huang
* @Last Modified time: 2018-04-12 15:13:44
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

type Location struct {
	Coordinates []float64 `json:"coordinates"`
}

type ArticleReq struct {
	SessionKey string    `json:"sessionkey"` //session_key
	AuthorId   string    `json:"authorid"`   //作者Id
	AuthorName string    `json:"authorname"` //作者名字
	AuthorHead string    `json:"authorhead"` //作者头像
	Loc        *Location `json:"loc"`        //写的位置
	Content    string    `json:"content"`    //内容
	Images     []string  `json:"images"`     //图像地址
	Anonymous  bool      `json:"anon"`       //是否匿名
}

// ============================================================================

func EditHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	if r.Method != "POST" {
		return
	}

	result, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	var req ArticleReq
	err := json.Unmarshal([]byte(result), &req)
	log.Info(req)

	if err != nil {
		log.Error("json Unmarshal error", result, err)
		w.Write([]byte(ErrEditFailed))
		return
	}

	if !CheckSessionKey(req.AuthorId, req.SessionKey) {
		log.Error("CheckSessionKey error", req.AuthorId, req.SessionKey)
		w.Write([]byte(ErrEditFailed))
		return
	}

	if len(req.Loc.Coordinates) != 2 {
		log.Error("Coordinates error", req.Loc.Coordinates)
		w.Write([]byte(ErrEditFailed))
		return
	}

	name, head := req.AuthorName, req.AuthorHead
	//匿名处理
	if req.Anonymous {
		name = GenRandName()
		head = AnonymousHead()
	}

	dbmgr.WriteArticle(&dbmgr.Articles{
		AuthorId:   req.AuthorId,
		AuthorName: name,
		AuthorHead: head,
		Loc: &dbmgr.Location{
			Type:        "Point",
			Coordinates: req.Loc.Coordinates,
		},
		Ts:        time.Now(),
		Content:   req.Content,
		Anonymous: req.Anonymous,
		Images:    req.Images,
	})

	w.Write([]byte(Success))

	//只保存发言的用户名和头像
	err = dbmgr.CenterUpdateUserNameHead(req.AuthorId, req.AuthorName, req.AuthorHead)
	if err != nil {
		log.Error("CenterUpdateUserInfo error", err)
		return
	}
}
