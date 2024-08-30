package scanx

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/goccy/go-json"
)

/**
 * @desc: 单条scan转换成map
 * @param rows 要转换的数据
 * @return {*}
 */
func PgRowToMap(rows *sql.Rows) map[string]interface{} {
	colTypes, _ := rows.ColumnTypes()                 // 列信息
	var rowParam = make([]interface{}, len(colTypes)) // 传入到 rows.Scan 的参数 数组
	var rowValue = make([]interface{}, len(colTypes)) // 接收数据一行列的数组

	for i, colType := range colTypes {
		rowValue[i] = reflect.New(colType.ScanType())           // 跟据数据库参数类型，创建默认值 和类型
		rowParam[i] = reflect.ValueOf(&rowValue[i]).Interface() // 跟据接收的数据的类型反射出值的地址
	}
	record := make(map[string]interface{})

	for rows.Next() {
		rows.Scan(rowParam...)
		for i, colType := range colTypes {
			switch rowValue[i].(type) {
			case []byte:
				if strings.HasSuffix(colType.Name(), "_ids") || strings.HasSuffix(colType.Name(), "_range") { // 多选ID类型
					var ids []uint
					json.Unmarshal(rowValue[i].([]byte), &ids)
					record[colType.Name()] = ids
				} else if strings.HasSuffix(colType.Name(), "_cascader") || strings.HasSuffix(colType.Name(), "_files") || strings.HasSuffix(colType.Name(), "_photos") || strings.HasSuffix(colType.Name(), "_checkbox") { // 多选数组字符串类型
					var rec []string
					json.Unmarshal(rowValue[i].([]byte), &rec)
					record[colType.Name()] = rec
				} else if strings.HasSuffix(colType.Name(), "_arr") { // 多选数组字符串类型
					var rec interface{}
					json.Unmarshal(rowValue[i].([]byte), &rec)
					record[colType.Name()] = rec
				} else {
					var rec map[string]interface{}
					json.Unmarshal(rowValue[i].([]byte), &rec)
					record[colType.Name()] = rec
				}
			case string:
				if strings.HasSuffix(colType.Name(), "_number") || strings.HasSuffix(colType.Name(), "_rate") || strings.HasSuffix(colType.Name(), "_slider") || strings.HasSuffix(colType.Name(), "_money") {
					val, _ := convertor.ToFloat(rowValue[i])
					record[colType.Name()] = val
				} else {
					record[colType.Name()] = rowValue[i]
				}

			default:
				record[colType.Name()] = rowValue[i]
			}
		}
	}
	return record
}

/**
 * @desc: 多条scan转换成map
 * @param undefined
 * @return {*}
 */
func PgRowsToMap(rows *sql.Rows) []map[string]interface{} {
	res := make([]map[string]interface{}, 0)          //  定义结果 map
	colTypes, _ := rows.ColumnTypes()                 // 列信息
	var rowParam = make([]interface{}, len(colTypes)) // 传入到 rows.Scan 的参数 数组
	var rowValue = make([]interface{}, len(colTypes)) // 接收数据一行列的数组

	for i, colType := range colTypes {
		rowValue[i] = reflect.New(colType.ScanType())           // 跟据数据库参数类型，创建默认值 和类型
		rowParam[i] = reflect.ValueOf(&rowValue[i]).Interface() // 跟据接收的数据的类型反射出值的地址
	}
	// 遍历
	for rows.Next() {
		rows.Scan(rowParam...) // 赋值到 rowValue 中
		record := make(map[string]interface{})
		for i, colType := range colTypes {
			switch rowValue[i].(type) {
			case []byte:
				if strings.HasSuffix(colType.Name(), "_ids") {
					var ids []uint
					json.Unmarshal(rowValue[i].([]byte), &ids)
					record[colType.Name()] = ids
				} else {
					var rec map[string]interface{}
					json.Unmarshal(rowValue[i].([]byte), &rec)
					record[colType.Name()] = rec
				}
			default:
				record[colType.Name()] = rowValue[i]
			}
		}
		res = append(res, record)
	}
	return res
}
