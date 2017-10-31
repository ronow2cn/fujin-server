/*
* @Author: huang
* @Date:   2017-10-24 14:21:05
* @Last Modified by:   huang
* @Last Modified time: 2017-10-31 17:48:37
 */
package controllers

import (
	"comm/config"
	"comm/dbmgr"
	"comm/logger"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"fujin/randname"
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

// 返回值距离
func EarthDistance(lng1, lat1, lng2, lat2 float64) string {
	radius := float64(6371000) // 6378137
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1

	dist := int(math.Acos(math.Sin(lat1)*math.Sin(lat2)+math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta)) * radius)

	if dist < 100 {
		return "小于100米"
	} else if dist < 1000 {
		return fmt.Sprintf("%d米", dist)
	} else if dist < 10000 {
		return fmt.Sprintf("%.1f千米", float32(dist)/1000)
	} else if dist < 10000000 {
		return fmt.Sprintf("%d千米", dist/1000)
	} else {
		return "十万八千里"
	}
}

//距今多久之前
func TimeGapStr(ts time.Time) string {
	gap := time.Since(ts).Minutes()

	if gap <= 0 {
		return "刚刚"
	} else if gap < 60 {
		return fmt.Sprintf("%d分钟前", int32(gap))
	} else if gap < 1440 {
		return fmt.Sprintf("%d小时前", int32(gap)/60)
	} else {
		_, m, d := ts.Date()
		return fmt.Sprintf("%d/%d", d, m)
	}

	return "long time ago"
}

//生成随机名字
func GenRandName() string {
	l := len(randname.RandName)
	r := rand_image.Intn(l)

	return fmt.Sprintf("%s%s", randname.RandNameFix, randname.RandName[r])
}

//匿名头像
func AnonymousHead() string {
	return config.Common.AnonHead
}
