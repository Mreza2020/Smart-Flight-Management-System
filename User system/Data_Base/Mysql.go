package Data_Base

import (
	"User_system/Error_Handling"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func DBMysql() (*sql.DB, string) {
	db, err := sql.Open("mysql", Error_Handling.LoadConfigFile().DBMysql.DB)
	if err != nil {
		zap.L().Error("not open mysql server", zap.String("err", err.Error()))
		return nil, "err"
	}
	if err = db.Ping(); err != nil {
		zap.L().Error("not ping", zap.String("err", err.Error()))
		return nil, "err"

	}
	zap.L().Info("Server started successfully")
	return db, ""
}
