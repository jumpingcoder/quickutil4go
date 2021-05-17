package dbutil

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jumpingcoder/quickutil4go/utils/logutil"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

var dbs map[string]*sqlx.DB = make(map[string]*sqlx.DB)

func Init(configs []interface{}, decryptHandler func(content string) string) bool {
	for _, config := range configs {
		configMap := config.(map[string]interface{})
		if configMap["DBName"] == nil || configMap["Driver"] == nil || configMap["Url"] == nil {
			logutil.Error("DBName、Driver、Url为必填项", nil)
			dbs = nil
			return false
		}
		dbname := configMap["DBName"].(string)
		db, err := sqlx.Open(configMap["Driver"].(string), decryptHandler(configMap["Url"].(string)))
		if err != nil {
			logutil.Error("数据库"+dbname+"初始化失败", err)
		}
		if configMap["MaxOpenConns"] != nil {
			db.SetMaxOpenConns(int(configMap["MaxOpenConns"].(float64)))
		}
		if configMap["MaxIdleConns"] != nil {
			db.SetMaxIdleConns(int(configMap["MaxIdleConns"].(float64)))
		}
		if configMap["ConnMaxIdleTime"] != nil {
			db.SetConnMaxIdleTime(time.Duration(time.Second.Nanoseconds() * int64(configMap["ConnMaxIdleTime"].(float64))))
		}
		if configMap["ConnMaxLifetime"] != nil {
			db.SetConnMaxLifetime(time.Duration(time.Second.Nanoseconds() * int64(configMap["ConnMaxLifetime"].(float64))))
		}
		err = db.Ping()
		if err != nil {
			logutil.Error("数据库"+dbname+"连接失败", err)
			dbs = nil
			return false
		}
		dbs[dbname] = db
	}
	return true
}

func DB(dbname string) *sqlx.DB {
	return dbs[dbname]
}

func QueryMap(dbname string, query string, args ...interface{}) []map[string]interface{} {
	stmt, err := dbs[dbname].Prepare(query)
	if err != nil {
		logutil.Error(nil, err)
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		logutil.Error(nil, err)
		return nil
	}
	var resultList []map[string]interface{}
	for rows.Next() {
		//字段列表
		columns, err := rows.Columns()
		if err != nil {
			logutil.Error(nil, err)
			return nil
		}
		//字段类型列表
		columnTypes, err := rows.ColumnTypes()
		//值列表
		values := make([]interface{}, len(columns))
		for i := range values {
			values[i] = new(interface{})
		}
		err = rows.Scan(values...)
		if err != nil {
			logutil.Error(nil, err)
			return nil
		}
		//组装
		result := make(map[string]interface{})
		for i, column := range columns {
			if strings.Index(columnTypes[i].DatabaseTypeName(), "VARCHAR") >= 0 {
				result[column] = string((*(values[i].(*interface{}))).([]uint8))
			} else {
				result[column] = *(values[i].(*interface{}))
			}
		}
		resultList = append(resultList, result)
	}
	logutil.Info(resultList, nil)
	return resultList
}
