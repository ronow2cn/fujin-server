/*
* @Author: huang
* @Date:   2018-05-08 11:16:02
* @Last Modified by:   huang
* @Last Modified time: 2018-05-08 11:46:25
 */
package controllers

import (
	"comm/dbmgr"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ============================================================================

type DelArticleReq struct {
	SessionKey string `json:"sessionkey"` //session_key
	Uid        string `json:"uid"`        //用户id
	ArticleId  string `json:"articleid"`  //删除文章id
}

// ============================================================================

func DelArticleHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	if r.Method != "POST" {
		return
	}

	result, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	var req DelArticleReq

	err := json.Unmarshal([]byte(result), &req)
	if err != nil {
		log.Error("json Unmarshal error", result, err)
		w.Write([]byte(ErrDelArticleFailed))
		return
	}

	if !CheckSessionKey(req.Uid, req.SessionKey) {
		log.Error("CheckSessionKey error", req.Uid, req.SessionKey)
		w.Write([]byte(ErrTokenError))
		return
	}

	err = dbmgr.CenterDelArticle(req.Uid, req.ArticleId)
	if err != nil {
		return
	}

	w.Write([]byte(Success))
}
