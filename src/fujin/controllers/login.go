/*
* @Author: huang
* @Date:   2017-10-25 15:08:47
* @Last Modified by:   huang
* @Last Modified time: 2017-10-25 16:32:56
 */
package controllers

import (
	"comm"
	"comm/config"
	"comm/dbmgr"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WeiXinAuthRes struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Expire     int32  `json:"expires_in"`
	Unionid    string `json:"unionid"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	code := r.PostFormValue("code")
	if len(code) == 0 {
		log.Error("code error", code)
		w.Write([]byte(ErrLoginFailed))
		return
	}

	url := fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		config.WeiXin.Code2SessionUrl, config.WeiXin.AppId, config.WeiXin.AppKey, code)
	log.Info("LoginHandler", url)

	// 用code去微信换取session_key
	ret, err := comm.HttpGetT(url, 5)
	if err != nil {
		log.Error("code to weixin auth error", code, err)
		w.Write([]byte(ErrLoginFailed))
		return
	}
	log.Info("weixin server res", ret)

	// 解析微信服务器返回的结果json
	var jret WeiXinAuthRes
	err = json.Unmarshal([]byte(ret), &jret)
	if err != nil {
		log.Error("json Unmarshal error", ret, err)
		w.Write([]byte(ErrLoginFailed))
		return
	}

	openidLen, sessLen := len(jret.OpenId), len(jret.SessionKey)
	if openidLen == 0 || sessLen == 0 {
		log.Error("openid len error", openidLen, sessLen)
		w.Write([]byte(ErrLoginFailed))
		return
	}

	// 将openid相关信息存入db
	expire := time.Now().Add(time.Duration(24) * time.Hour)
	err = dbmgr.CenterUpdateUserInfo(jret.OpenId, jret.SessionKey, time.Now().Add(time.Duration(24)*time.Hour), jret.Unionid)
	if err != nil {
		log.Error("CenterUpdateUserInfo error", err)
		w.Write([]byte(ErrLoginFailed))
		return
	}

	//返回给用户信息
	res := fmt.Sprintf("{\"openid\":\"%s\",\"session_key\":\"%s\",\"expires_in\":%d}",
		jret.OpenId, jret.SessionKey[:(sessLen/2)], expire.Unix())

	//Success
	w.Write([]byte(res))
}
