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
	LogLevel    string `json:"logLevel"`
	DBCenter    string `json:"dbCenter"`
	Port        int32  `json:"port"`
	Images      string `json:"images"`
	ImagesUrl   string `json:"imagesurl"`
	GmToken     string `json:"gmtoken"`
	MaxDistance int32  `json:"maxdistance"`
	AnonHead    string `json:"anonhead"`
}

type weixinT struct {
	AppId           string `json:"appid"`
	AppKey          string `json:"appkey"`
	Code2SessionUrl string `json:"code2sessurl"`
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
