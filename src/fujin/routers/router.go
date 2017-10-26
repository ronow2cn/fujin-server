/*
* @Author: huang
* @Date:   2017-10-24 11:07:21
* @Last Modified by:   huang
* @Last Modified time: 2017-10-26 15:42:10
 */
package routers

import (
	"comm/config"
	"comm/logger"
	"fmt"
	"fujin/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

// ============================================================================

var log = logger.DefaultLogger

// ============================================================================

func httpServer(r *mux.Router) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Common.Port), r)
	if err != nil {
		log.Error("http service exited:", err)
	}
}

func httpsServer(r *mux.Router) {
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", config.Common.Port),
		"www.esiyou.com.pem", "www.esiyou.com.key", r)
	if err != nil {
		log.Fatal("ListenAndServeTLS:", err.Error())
	}
}

// ============================================================================

func Routers() {
	go func() {
		r := mux.NewRouter()
		r.HandleFunc("/", controllers.HelloHandler).Methods("GET")
		r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
		r.HandleFunc("/uploadfile", controllers.UploadFileHandler).Methods("POST")
		r.HandleFunc("/edit", controllers.EditHandler).Methods("POST")
		r.HandleFunc("/found", controllers.FoundHandler).Methods("POST")
		r.HandleFunc("/callback", controllers.CallBackHandler).Methods("POST")
		r.HandleFunc("/comment", controllers.CommentHandler).Methods("POST")

		httpServer(r)
	}()
}
