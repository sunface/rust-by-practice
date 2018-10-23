package g

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

func InitMysql(acc, pw, addr, port, database string) {
	var err error
	// 初始化mysql连接
	sqlConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", acc, pw, addr, port, database)
	DB, err = sqlx.Open("mysql", sqlConn)
	if err != nil {
		L.Fatal("init mysql error", zap.Error(err))
	}

	// 测试db是否正常
	err = DB.Ping()
	if err != nil {
		L.Fatal("init mysql, ping error", zap.Error(err))
	}
}
