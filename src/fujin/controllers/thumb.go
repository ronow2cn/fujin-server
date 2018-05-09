/*
* @Author: huang
* @Date:   2018-05-09 10:48:35
* @Last Modified by:   huang
* @Last Modified time: 2018-05-09 16:27:08
 */
package controllers

import (
	"comm/dbmgr"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ============================================================================

type ThumbReq struct {
	SessionKey string `json:"sessionkey"` //session_key
	Uid        string `json:"uid"`        //用户id
	Name       string `json:"name"`       //名字
	Head       string `json:"head"`       //头像
	ThumbType  string `json:"thumbtype"`  //点赞类型
	ArticleId  string `json:"articleid"`  //点赞文章id
	CommentId  string `json:"commentid"`  //点赞评论id
}

// ============================================================================

func ThumbHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	if r.Method != "POST" {
		return
	}

	result, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	var req ThumbReq

	err := json.Unmarshal([]byte(result), &req)
	if err != nil {
		log.Error("json Unmarshal error", result, err)
		w.Write([]byte(ErrThumbFailed))
		return
	}

	if !CheckSessionKey(req.Uid, req.SessionKey) {
		log.Error("CheckSessionKey error", req.Uid, req.SessionKey)
		w.Write([]byte(ErrTokenError))
		return
	}

	if req.ThumbType == "comment" {
		err = dbmgr.CommentThumbAdd(req.Uid, req.ArticleId, req.CommentId)
		if err != nil {
			return
		}
	} else {
		err = dbmgr.ArticleThumbAdd(req.Uid, req.ArticleId)
		if err != nil {
			return
		}
	}

	w.Write([]byte(Success))

	//只保存发言的用户名和头像
	err = dbmgr.CenterUpdateUserNameHead(req.Uid, req.Name, req.Head)
	if err != nil {
		log.Error("CenterUpdateUserInfo error", err)
		return
	}
}
