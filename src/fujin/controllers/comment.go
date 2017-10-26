/*
* @Author: huang
* @Date:   2017-10-26 15:22:38
* @Last Modified by:   huang
* @Last Modified time: 2017-10-26 15:51:11
 */
package controllers

import (
	"comm/dbmgr"
	"net/http"
	"time"
)

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
	c := &dbmgr.CommentOne{
		CUid:  "comment1",
		CName: "name",
		CHead: "head",
		Loc: &dbmgr.Location{
			Type:        "Point",
			Coordinates: []float64{104.066541, 30.572269},
		},
		Ts:        time.Now(),
		Content:   "content",
		Anonymous: false,
	}

	dbmgr.WriteComment("15089300623", c)
}
