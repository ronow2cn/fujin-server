/*
* @Author: huang
* @Date:   2017-10-26 11:57:22
* @Last Modified by:   huang
* @Last Modified time: 2017-10-26 12:08:50
 */

package controllers

import (
	"comm/dbmgr"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CallBackReq struct {
	Uid     string `json:"uid"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func CallBackHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	if r.Method != "POST" {
		return
	}

	result, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	var req CallBackReq
	err := json.Unmarshal([]byte(result), &req)
	log.Info(req)

	if err != nil {
		log.Error("json Unmarshal error", result, err)
		w.Write([]byte(ErrCallBackFailed))
		return
	}

	dbmgr.InsertCallBack(req.Uid, req.Name, req.Content)

	w.Write([]byte(Success))
}
