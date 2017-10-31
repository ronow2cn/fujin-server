/*
* @Author: huang
* @Date:   2017-10-24 11:09:54
* @Last Modified by:   huang
* @Last Modified time: 2017-10-31 17:31:11
 */
package controllers

import (
	"io"
	"net/http"
	//"time"
)

// ============================================================================

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "hello world!")

}
