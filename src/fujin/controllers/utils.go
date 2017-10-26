/*
* @Author: huang
* @Date:   2017-10-24 14:21:05
* @Last Modified by:   huang
* @Last Modified time: 2017-10-26 14:45:01
 */
package controllers

import (
	"comm/config"
	"comm/dbmgr"
	"comm/logger"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

// ============================================================================
// imageid[:]    --> BA2871C220171020
// imageid[0-8]  --> BA2871C2  [是8个随机的字符,16进制的字符类型]
// imageid[8:]   --> 20171020  [年月日时间]
// ============================================================================

var log = logger.DefaultLogger

var rand_image = rand.New(rand.NewSource(time.Now().Unix()))

// ============================================================================

func MakeImageID() string {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, rand_image.Uint32())

	y, m, d := time.Now().Date()

	return fmt.Sprintf("%s%d%d%d", strings.ToUpper(hex.EncodeToString(buf)), y, m, d)
}

func ImageID2Path(imageid, filetype string) string {
	if filetype == "image/png" {
		return fmt.Sprintf("%s/%s/%s.png", config.Common.Images, imageid[8:], imageid[0:8])
	} else {
		return fmt.Sprintf("%s/%s/%s.jpg", config.Common.Images, imageid[8:], imageid[0:8])
	}
}

func ImageID2Url(imageid, filetype string) string {
	if filetype == "image/png" {
		return fmt.Sprintf("%s/%s/%s.png", config.Common.ImagesUrl, imageid[8:], imageid[0:8])
	} else {
		return fmt.Sprintf("%s/%s/%s.jpg", config.Common.ImagesUrl, imageid[8:], imageid[0:8])
	}
}

func FileExist(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return false
	} else {
		return true
	}
}

func BuildTree(imageid string) error {
	return os.MkdirAll(fmt.Sprintf("%s/%s", config.Common.Images, imageid[8:]), 0755)
}

//sessionKey 是要验证的验证码
func CheckSessionKey(uid, sessionKey string) bool {
	if sessionKey == config.Common.GmToken {
		return true
	}

	user := dbmgr.CenterGetUserInfo(uid)
	if user == nil {
		return false
	}

	l := len(user.SessionKey)
	if l > 0 && (user.SessionKey[:(l/2)] == sessionKey) {
		return true
	}

	return false
}

// 返回值的单位为米
func EarthDistance(lng1, lat1, lng2, lat2 float64) float64 {
	radius := float64(6371000) // 6378137
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}
