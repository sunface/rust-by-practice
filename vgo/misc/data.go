package misc

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mafanr/g"
	"go.uber.org/zap"
)

func InitMysql() {
	var err error

	// 初始化mysql连接
	sqlConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", Conf.Mysql.Acc, Conf.Mysql.Pw,
		Conf.Mysql.Addr, Conf.Mysql.Port, Conf.Mysql.Database)
	g.DB, err = sqlx.Open("mysql", sqlConn)
	if err != nil {
		g.L.Fatal("init mysql error", zap.Error(err))
	}

	// 测试db是否正常
	err = g.DB.Ping()
	if err != nil {
		g.L.Fatal("init mysql, ping error", zap.Error(err))
	}
}
