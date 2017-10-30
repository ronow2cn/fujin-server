package dbmgr

import (
	"comm/config"
	"comm/db"
)

// ============================================================================

const (
	// center
	CTableUsers    = "users"
	CTableArticles = "articles"
	CTableSeqId    = "seqid"
	CTableCallback = "callback"
	CTableComments = "comments"
)

// ============================================================================

var (
	DBCenter *db.Database
)

// ============================================================================

func Open() {
	// 初始化 中心 数据库
	if DBCenter == nil {
		DBCenter = db.NewDatabase()
		DBCenter.Open(config.Common.DBCenter, false)
	}

	CenterCreateSeqId()

	DBCenter.CreateIndex(CTableUsers, "uid", []string{"_id"}, true)

	DBCenter.CreateIndex(CTableArticles, "loc", []string{"$2dsphere:loc"}, false)
	DBCenter.CreateIndex(CTableArticles, "authorid", []string{"authorid"}, false)

}

func Close() {
	DBCenter.Close()
}

// ============================================================================
