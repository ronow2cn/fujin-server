/*
* @Author: huang
* @Date:   2017-10-25 17:07:01
* @Last Modified by:   huang
* @Last Modified time: 2018-05-09 14:24:12
 */
package dbmgr

import (
	"comm/config"
	"comm/db"
	"time"
)

// ============================================================================

type Location struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type Articles struct {
	Id         string    `bson:"_id"`        //Id
	AuthorId   string    `bson:"authorid"`   //作者Id
	AuthorName string    `bson:"authorname"` //作者名字
	AuthorHead string    `bson:"authorhead"` //作者头像
	Loc        *Location `bson:"loc"`        //写的位置
	Ts         time.Time `bson:"ts"`         //写时间
	Content    string    `bson:"content"`    //内容
	Images     []string  `bson:"images"`     //图像地址
	Anonymous  bool      `bson:"anon"`       //是否匿名
	IsDelete   bool      `bson:"isdelete"`   //是否已经删除
}

// ============================================================================

func WriteArticle(article *Articles) {
	var obj Articles
	obj.Id = GenArticleId()
	obj.AuthorId = article.AuthorId
	obj.AuthorName = article.AuthorName
	obj.AuthorHead = article.AuthorHead
	obj.Loc = article.Loc
	obj.Ts = article.Ts
	obj.Content = article.Content

	if len(article.Images) > 0 {
		obj.Images = append(obj.Images, article.Images...)
	}

	obj.Anonymous = article.Anonymous

	err := DBCenter.Insert(CTableArticles, &obj)
	if err != nil {
		log.Error("WriteArticle error", err)
	}

}

func GetArticlesById(id string) *Articles {
	var obj Articles

	err := DBCenter.GetObjectByCond(
		CTableArticles,
		db.M{
			"_id": id,
		},
		&obj,
	)

	if err == nil {
		return &obj
	} else {
		// failed
		return nil
	}
}

func GetArticlesByAuthorId(authorid string) (ret []*Articles) {

	err := DBCenter.GetAllObjectsByCond(
		CTableArticles,
		db.M{
			"authorid": authorid,
			"isdelete": false,
		},
		&ret,
	)

	if err != nil {
		log.Error("GetArticlesAuthorId", authorid, err)
	}

	return
}

func GetArticlesByLocation(longitude, latitude float64, distance int32) (ret []*Articles) {
	Arr := []float64{longitude, latitude}
	if distance <= 0 || distance > config.Common.MaxDistance {
		distance = config.Common.MaxDistance
	}

	err := DBCenter.GetAllObjectsByCond(
		CTableArticles,
		db.M{
			"loc": db.M{
				"$near": db.M{
					"$geometry": db.M{
						"type":        "Point",
						"coordinates": Arr,
					},
					"$maxDistance": distance,
				},
			},
			"isdelete": false,
		},
		&ret,
	)

	if err != nil {
		log.Error("GetArticlesAuthorId", longitude, latitude, err)
	}

	return
}

func GetArticlesByAuthorIdLimit(authorid string, skip, limit int) (ret []*Articles) {

	err := DBCenter.GetObjectsBySkipLimited(
		CTableArticles,
		db.M{
			"authorid": authorid,
			"isdelete": false,
		},
		skip,
		limit,
		&ret,
	)

	if err != nil {
		log.Error("GetArticlesAuthorId", authorid, err)
	}

	return
}

func GetArticlesByLocationByLimit(longitude, latitude float64, distance int32, skip, limit int) (ret []*Articles) {
	Arr := []float64{longitude, latitude}
	if distance <= 0 || distance > config.Common.MaxDistance {
		distance = config.Common.MaxDistance
	}

	err := DBCenter.GetObjectsBySkipLimited(
		CTableArticles,
		db.M{
			"loc": db.M{
				"$near": db.M{
					"$geometry": db.M{
						"type":        "Point",
						"coordinates": Arr,
					},
					"$maxDistance": distance,
				},
			},
			"isdelete": false,
		},
		skip,
		limit,
		&ret,
	)

	if err != nil {
		log.Error("GetArticlesByLocationByLimit", longitude, latitude, skip, limit, err)
	}

	return
}

func GetArticlesByLimit(skip, limit int) (ret []*Articles) {

	err := DBCenter.GetObjectsBySkipLimited(
		CTableArticles,
		db.M{
			"isdelete": false,
		},
		skip,
		limit,
		&ret,
	)

	if err != nil {
		log.Error("GetArticlesByLocationByLimit", skip, limit, err)
	}

	return
}

// 删除文章
func CenterDelArticle(authorid string, articleid string) error {
	err := DBCenter.UpdateByCond(
		CTableArticles,
		db.M{
			"_id":      articleid,
			"authorid": authorid,
		},
		db.M{
			"$set": db.M{
				"isdelete": true,
			},
		},
	)

	if err != nil {
		log.Warning("CenterDelArticle failed:", err)
	}

	return err
}
