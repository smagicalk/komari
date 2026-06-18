package flags

import (
	"fmt"
	"log/slog"
	"strings"
)

const (
	DatabaseTypeSQLite   = "sqlite"
	DatabaseTypeMySQL    = "mysql"
	DatabaseTypePostgres = "postgres"
)

var (
	// 数据库配置
	DatabaseType string // 数据库类型：sqlite, mysql, postgres
	DatabaseFile string // SQLite数据库文件路径
	DatabaseHost string // MySQL/PostgreSQL 数据库主机地址
	DatabasePort string // MySQL/PostgreSQL 数据库端口
	DatabaseUser string // MySQL/PostgreSQL 数据库用户名
	DatabasePass string // MySQL/PostgreSQL 数据库密码
	DatabaseName string // MySQL/PostgreSQL 数据库名称

	Listen   string
	LogLevel string // 日志级别：debug, info, warn, error
)

func NormalizeDatabaseType(databaseType string) string {
	databaseType = strings.ToLower(strings.TrimSpace(databaseType))
	if databaseType == "" {
		return DatabaseTypeSQLite
	}
	return databaseType
}

func ApplyDatabaseTypeNormalization() string {
	DatabaseType = NormalizeDatabaseType(DatabaseType)
	return DatabaseType
}

func IsSQLite() bool {
	return NormalizeDatabaseType(DatabaseType) == DatabaseTypeSQLite
}

func SupportedDatabaseTypes() string {
	return DatabaseTypeSQLite + ", " + DatabaseTypeMySQL + ", " + DatabaseTypePostgres
}

func NormalizeLogLevel(level string) string {
	level = strings.ToLower(strings.TrimSpace(level))
	if level == "" {
		return ""
	}
	return level
}

func ParseLogLevel(level string) (slog.Level, error) {
	switch NormalizeLogLevel(level) {
	case "debug":
		return slog.LevelDebug, nil
	case "", "info":
		return slog.LevelInfo, nil
	case "warn", "warning":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("unsupported log level: %s", level)
	}
}
