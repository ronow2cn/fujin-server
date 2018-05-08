/*
* @Author: huang
* @Date:   2018-05-08 14:47:51
* @Last Modified by:   huang
* @Last Modified time: 2018-05-08 15:01:58
 */
package controllers

import (
	"comm/dbmgr"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ============================================================================

type DelCommentReq struct {
	SessionKey string `json:"sessionkey"` //session_key
	Uid        string `json:"uid"`        //用户id
	ArticleId  string `json:"articleid"`  //删除文章id
	CommentId  string `json:"commentid"`  //删除id
}

// ============================================================================
// ============================================================================

func DelCommentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	if r.Method != "POST" {
		return
	}

	result, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	var req DelCommentReq

	err := json.Unmarshal([]byte(result), &req)
	if err != nil {
		log.Error("json Unmarshal error", result, err)
		w.Write([]byte(ErrDelCommentFailed))
		return
	}

	if !CheckSessionKey(req.Uid, req.SessionKey) {
		log.Error("CheckSessionKey error", req.Uid, req.SessionKey)
		w.Write([]byte(ErrTokenError))
		return
	}

	err = dbmgr.CenterDelComment(req.Uid, req.ArticleId, req.CommentId)
	if err != nil {
		return
	}

	w.Write([]byte(Success))
}
