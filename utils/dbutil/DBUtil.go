package dbutil

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jumpingcoder/quickutil4go/utils/cryptoutil"
	"github.com/jumpingcoder/quickutil4go/utils/logutil"
	_ "github.com/lib/pq"
	"strings"
	"time"
)

var dbs map[string]*sqlx.DB = make(map[string]*sqlx.DB)

func InitFromConfig(configs []interface{}, decryptKey string, decryptHandler func(content string, decryptKey string) string) bool {
	for _, config := range configs {
		configMap := config.(map[string]interface{})
		if configMap["DBName"] == nil || configMap["Driver"] == nil || configMap["Url"] == nil {
			logutil.Error("DBName、Driver、Url为必填项", nil)
			dbs = nil
			return false
		}
		dbname := configMap["DBName"].(string)
		driver := configMap["Driver"].(string)
		url := decryptHandler(configMap["Url"].(string), decryptKey)
		maxOpenConns := 0
		maxIdleConns := -1
		connMaxIdleTime := time.Duration(0)
		connMaxLifetime := time.Duration(0)
		if configMap["MaxOpenConns"] != nil {
			maxOpenConns = int(configMap["MaxOpenConns"].(float64))
		}
		if configMap["MaxIdleConns"] != nil {
			maxIdleConns = int(configMap["MaxIdleConns"].(float64))
		}
		if configMap["ConnMaxIdleTime"] != nil {
			connMaxIdleTime = time.Duration(time.Second.Nanoseconds() * int64(configMap["ConnMaxIdleTime"].(float64)))
		}
		if configMap["ConnMaxLifetime"] != nil {
			connMaxLifetime = time.Duration(time.Second.Nanoseconds() * int64(configMap["ConnMaxLifetime"].(float64)))
		}
		AddDB(dbname, driver, url, maxOpenConns, maxIdleConns, connMaxIdleTime, connMaxLifetime)
	}
	return true
}

func AddDB(dbname string, driver string, url string, maxOpenConns int, maxIdleConns int, connMaxIdleTime time.Duration, connMaxLifetime time.Duration) {
	if dbs[dbname] != nil {
		logutil.Warn("Database "+dbname+" already exists", nil)
	}
	db, err := sqlx.Open(driver, url)
	if err != nil {
		logutil.Error("Database "+dbname+" open failed", err)
	}
	db = db.Unsafe()
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime)
	db.SetConnMaxLifetime(connMaxLifetime)
	err = db.Ping()
	if err != nil {
		logutil.Error("Database "+dbname+" connection failed", err)
		dbs = nil
	}
	dbs[dbname] = db
}

func GetDB(dbname string) *sqlx.DB {
	return dbs[dbname]
}


func DefaultDecryptHandler(content string, decryptKey string) string {
	start := strings.Index(content, "ENC(")
	if start < 1 {
		return content
	}
	end := strings.Index(content[start:len(content)], ")")
	password := content[start+4 : start+end]
	decrypted := string(cryptoutil.AESCBCDecrypt(cryptoutil.Base64Decrypt(password), []byte(decryptKey), make([]byte, 16)))
	newContent := content[0:start] + decrypted + content[start+end+1:len(content)]
	return newContent
}

//
//func QueryObject(dbname string, target interface{}, query string, args ...interface{}) []interface{} {
//	rows, err := dbs[dbname].Queryx(query, args...)
//	if err != nil {
//		logutil.Error(nil, err)
//		return nil
//	}
//	var resultList []interface{}
//	for rows.Next() {
//		var target entity.TestDO
//		err := rows.StructScan(&target)
//		if err != nil {
//			logutil.Error(nil, err)
//		}
//		resultList = append(resultList, target)
//	}
//	return resultList
//}

func QueryMap(dbname string, query string, args ...interface{}) []map[string]interface{} {
	//查询
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
	//驱动
	driverName := dbs[dbname].DriverName()
	//字段列表
	columnNames, err := rows.Columns()
	if err != nil {
		logutil.Error(nil, err)
		return nil
	}
	//字段类型列表
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		logutil.Error(nil, err)
		return nil
	}
	var resultList []map[string]interface{}
	//值列表
	for rows.Next() {
		values := make([]interface{}, len(columnNames))
		for i := range values {
			values[i] = new(interface{})
		}
		err = rows.Scan(values...)
		if err != nil {
			logutil.Error(nil, err)
			return nil
		}
		var result map[string]interface{}
		switch driverName {
		case "mysql":
			result = mysqlMapping(columnNames, columnTypes, values)
			break
		case "postgres":
			result = postgreMapping(columnNames, columnTypes, values)
			break
		default:
			result = defaultMapping(columnNames, columnTypes, values)
			break
		}
		resultList = append(resultList, result)
	}
	return resultList
}

func mysqlMapping(columnNames []string, columnTypes []*sql.ColumnType, values []interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for i, columnName := range columnNames {
		if *values[i].(*interface{}) == nil {
			result[columnName] = nil
			continue
		}
		if strings.Index(columnTypes[i].DatabaseTypeName(), "VARCHAR") == 0 {
			result[columnName] = string((*(values[i].(*interface{}))).([]uint8))
		} else if strings.Index(columnTypes[i].DatabaseTypeName(), "TEXT") == 0 {
			result[columnName] = string((*(values[i].(*interface{}))).([]uint8))
		} else if strings.Index(columnTypes[i].DatabaseTypeName(), "CHAR") == 0 {
			result[columnName] = string((*(values[i].(*interface{}))).([]uint8))
		} else if strings.Index(columnTypes[i].DatabaseTypeName(), "JSON") == 0 {
			content := make(map[string]interface{})
			json.Unmarshal((*(values[i].(*interface{}))).([]uint8), &content)
			result[columnName] = content
		} else {
			result[columnName] = *(values[i].(*interface{}))
		}
	}
	return result
}

func postgreMapping(columnNames []string, columnTypes []*sql.ColumnType, values []interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for i, columnName := range columnNames {
		if *values[i].(*interface{}) == nil {
			result[columnName] = nil
			continue
		}
		if strings.Index(columnTypes[i].DatabaseTypeName(), "BPCHAR") == 0 {
			result[columnName] = string((*(values[i].(*interface{}))).([]uint8))
		} else if strings.Index(columnTypes[i].DatabaseTypeName(), "JSON") == 0 {
			content := make(map[string]interface{})
			json.Unmarshal((*(values[i].(*interface{}))).([]uint8), &content)
			result[columnName] = content
		} else {
			result[columnName] = *(values[i].(*interface{}))
		}
	}
	return result
}

func defaultMapping(columnNames []string, columnTypes []*sql.ColumnType, values []interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for i, columnName := range columnNames {
		result[columnName] = *(values[i].(*interface{}))
	}
	return result
}
