/*
* @Author: huang
* @Date:   2017-10-24 11:07:21
* @Last Modified by:   huang
* @Last Modified time: 2018-05-09 10:57:35
 */
package routers

import (
	"comm/config"
	"comm/logger"
	"fmt"
	"fujin/controllers"
	"net/http"
)

// ============================================================================

var log = logger.DefaultLogger

// ============================================================================

func httpServer() {
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Common.Port), nil)
	if err != nil {
		log.Error("http service exited:", err)
	}
}

func httpsServer() {
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", config.Common.Port),
		"www.esiyou.com.pem", "www.esiyou.com.key", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS:", err.Error())
	}
}

// ============================================================================

func Routers() {
	go func() {
		http.HandleFunc("/", controllers.HelloHandler)
		http.HandleFunc("/login", controllers.LoginHandler)
		http.HandleFunc("/uploadfile", controllers.UploadFileHandler)
		http.HandleFunc("/edit", controllers.EditHandler)
		http.HandleFunc("/found", controllers.FoundHandler)
		http.HandleFunc("/callback", controllers.CallBackHandler)
		http.HandleFunc("/comment", controllers.CommentHandler)
		http.HandleFunc("/getcomment", controllers.GetCommentHandler)
		http.HandleFunc("/delarticle", controllers.DelArticleHandler)
		http.HandleFunc("/delcomment", controllers.DelCommentHandler)
		http.HandleFunc("/thumb", controllers.ThumbHandler)

		httpServer()
	}()
}
