package flags

import "strings"

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

	Listen string
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
