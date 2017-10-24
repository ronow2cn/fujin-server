/*
* @Author: huang
* @Date:   2017-10-24 15:44:39
* @Last Modified by:   huang
* @Last Modified time: 2017-10-24 17:59:18
 */
package controllers

import (
	"io"
	"net/http"
	"os"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	//上传参数为uploadfile
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		log.Error(err)
		w.Write([]byte(ErrUploadFileFailed))
		return
	}
	defer file.Close()
	//检测文件类型
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		log.Error(err)
		w.Write([]byte(ErrUploadFileFailed))
		return
	}
	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpg" {
		log.Info(filetype)
		w.Write([]byte(ErrUploadFileFailed))
		return
	}
	//文件指针
	log.Info(filetype)
	if _, err = file.Seek(0, 0); err != nil {
		log.Error(err)
	}

	//随机生成一个不存在的fileid
	var imgid string
	for {
		imgid = MakeImageID()
		if !FileExist(ImageID2Path(imgid, filetype)) {
			break
		}
	}

	//创建整棵存储树
	if err = BuildTree(imgid); err != nil {
		log.Error(err)
	}

	f, err := os.OpenFile(ImageID2Path(imgid, filetype), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Error(err)
		w.Write([]byte(ErrUploadFileFailed))
		return
	}
	defer f.Close()

	io.Copy(f, file)
	w.Write([]byte(ImageID2Url(imgid, filetype)))
}
