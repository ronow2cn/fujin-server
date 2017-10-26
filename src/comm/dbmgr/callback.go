/*
* @Author: huang
* @Date:   2017-10-26 12:02:21
* @Last Modified by:   huang
* @Last Modified time: 2017-10-26 12:07:56
 */
package dbmgr

type CallBack struct {
	Uid     string `bson:"uid"`
	Name    string `bson:"name"`
	Content string `bson:"content"`
}

func InsertCallBack(uid, name, content string) {
	var obj CallBack
	obj.Uid = uid
	obj.Name = name
	obj.Content = content

	// save to db
	err := DBCenter.Insert(CTableCallback, &obj)
	if err != nil {
		log.Error("create user failed:", uid, err)
	}
}
