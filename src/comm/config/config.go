package config

import (
	"comm"
	"encoding/json"
	"io/ioutil"
	"strings"
)

// ============================================================================

type commonT struct {
	Version     string `json:"version"`
	VerMajor    string `json:"-"`
	VerMinor    string `json:"-"`
	VerBuild    string `json:"-"`
	LogLevel    string `json:"logLevel"`    //日志等级
	DBCenter    string `json:"dbCenter"`    //db地址
	Port        int32  `json:"port"`        //监听端口
	Images      string `json:"images"`      //上传图片在服务器的路径
	ImagesUrl   string `json:"imagesurl"`   //上传图片的地址前缀
	GmToken     string `json:"gmtoken"`     //管理员sessionkey
	MaxDistance int32  `json:"maxdistance"` //显示最近几米的消息
	AnonHead    string `json:"anonhead"`    //匿名头像地址
	PerReqNum   int32  `json:"perreqnum"`   //每次返回查询的数量条数
}

type weixinT struct {
	AppId           string `json:"appid"`
	AppKey          string `json:"appkey"`
	Code2SessionUrl string `json:"code2sessurl"` //微信code换sessionkey的地址
}

// ============================================================================

type configT struct {
	Common *commonT `json:"common"`
	WeiXin *weixinT `json:"weixin"`
}

// ============================================================================

var (
	Common *commonT
	WeiXin *weixinT
)

// ============================================================================

func Parse(fn string, argServer string) {
	var conf configT

	// read file
	d, err := ioutil.ReadFile(fn)
	if err != nil {
		comm.Panic("open config file failed:", err)
	}

	// parse
	err = json.Unmarshal(d, &conf)
	if err != nil {
		comm.Panic("parse config file failed:", err)
	}

	parseVersion(&conf)

	// set variables
	Common = conf.Common
	WeiXin = conf.WeiXin

}

// ============================================================================

func parseVersion(conf *configT) {
	arr := strings.Split(conf.Common.Version, ".")
	if len(arr) < 3 {
		comm.Panic("invalid version:", conf.Common.Version)
	}

	conf.Common.VerMajor = arr[0]
	conf.Common.VerMinor = arr[1]
	conf.Common.VerBuild = arr[2]
}
