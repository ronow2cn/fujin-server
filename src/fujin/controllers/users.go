/*
* @Author: huang
* @Date:   2017-10-24 11:09:54
* @Last Modified by:   huang
* @Last Modified time: 2017-10-25 19:32:48
 */
package controllers

import (
	"comm/dbmgr"
	"io"
	"net/http"
	//"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	a := dbmgr.GetArticlesByAuthorId("authorid")

	for _, v := range a {
		log.Info(v)
	}

	io.WriteString(w, "hello world!")
	/*
		a := &dbmgr.Articles{
			AuthorId:   "authorid",
			AuthorName: "authorname",
			AuthorHead: "https://www.esiyou.com/wf/images/20171025/1F2C3121.png",
			Loc: &dbmgr.Location{
				Type:        "Point",
				Coordinates: []float64{104.066541, 30.572269},
			},
			Ts:      time.Now(),
			Content: "content",
			Anonymous: false,
			Images:  []string{"https://www.esiyou.com/wf/images/20171025/1F2C3121.png", "https://www.esiyou.com/wf/images/20171025/1F2C3121.png"},
		}

		dbmgr.WriteArticle(a)
	*/

}
