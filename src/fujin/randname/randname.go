/*
* @Author: huang
* @Date:   2017-10-31 16:51:05
* @Last Modified by:   huang
* @Last Modified time: 2017-10-31 17:12:29
 */
package randname

import (
	"comm/logger"
	"io/ioutil"
	"strings"
)

// ============================================================================

var log = logger.DefaultLogger

var RandNameFix = "隐匿的"

var RandName = []string{"蔷薇", "布袋"}

// ============================================================================

func Load(fn string) error {
	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(buf), ",") {
		line = strings.Trim(line, " \r")
		if line != "" {
			RandName = append(RandName, line)
		}
	}

	return nil
}
