package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

// TestConnection 測試資料庫連線
func TestConnection(db *Database) *ConnectionTestResult {
	if db == nil {
		return &ConnectionTestResult{
			Success: false,
			Message: "Database configuration is nil",
			Latency: 0,
		}
	}

	var dsn string
	var driverName string

	// 根據資料庫類型建立連線字串
	switch db.Type {
	case "mssql":
		driverName = "sqlserver"
		dsn = fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s",
			db.Host, db.Port, db.Username, db.Password, db.Database)
	case "mysql":
		driverName = "mysql"
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			db.Username, db.Password, db.Host, db.Port, db.Database)
	default:
		return &ConnectionTestResult{
			Success: false,
			Message: fmt.Sprintf("Unsupported database type: %s", db.Type),
			Latency: 0,
		}
	}

	// 記錄開始時間
	startTime := time.Now()

	// 嘗試連線
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		return &ConnectionTestResult{
			Success: false,
			Message: fmt.Sprintf("Failed to open connection: %v", err),
			Latency: 0,
		}
	}
	defer conn.Close()

	// Ping 測試
	if err := conn.Ping(); err != nil {
		return &ConnectionTestResult{
			Success: false,
			Message: fmt.Sprintf("Failed to ping database: %v", err),
			Latency: 0,
		}
	}

	// 計算延遲
	latency := time.Since(startTime).Milliseconds()

	return &ConnectionTestResult{
		Success: true,
		Message: "Connection successful",
		Latency: int(latency),
	}
}
