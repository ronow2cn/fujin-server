package dbmgr

import (
	"comm/config"
	"comm/db"
)

// ============================================================================

const (
	// center
	CTableUsers = "users"
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

	DBCenter.CreateIndex(CTableUsers, "idx_users", []string{"uid"}, true)
}

func Close() {
	DBCenter.Close()
}

// ============================================================================
