/*
* @Author: huang
* @Date:   2018-05-09 17:29:58
* @Last Modified by:   huang
* @Last Modified time: 2018-05-09 17:42:24
 */
package controllers

import (
	"comm/config"
	"encoding/json"
	"net/http"
)

// ============================================================================

type IndexImagesRes struct {
	Images []string `json:"indeximages"`
}

// ============================================================================

func IndexImagesHandler(w http.ResponseWriter, r *http.Request) {
	res := &IndexImagesRes{}

	for _, v := range config.Common.IndexImages {
		res.Images = append(res.Images, v)
	}

	if len(res.Images) == 0 {
		return
	}

	b, err := json.Marshal(res)
	if err != nil {
		log.Error("json.Marshal(res) error")
		return
	}

	w.Write([]byte(string(b)))
}
