/*
* @Author: huang
* @Date:   2017-10-24 11:09:54
* @Last Modified by:   huang
* @Last Modified time: 2017-10-24 11:18:03
 */
package controllers

import (
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!")
}
