/*
 * @Author: Coller
 * @Date: 2022-04-29 17:16:35
 * @LastEditTime: 2024-04-21 14:48:59
 * @Desc: 数据转换
 */
package scanx

import (
	"database/sql"
	"reflect"
)

/**
 * @desc: 单条scan转换成map
 * @param rows 要转换的数据
 * @return {*}
 */
func MyRowToMap(rows *sql.Rows) map[string]interface{} {
	colTypes, _ := rows.ColumnTypes()                 // 列信息
	var rowParam = make([]interface{}, len(colTypes)) // 传入到 rows.Scan 的参数 数组
	var rowValue = make([]interface{}, len(colTypes)) // 接收数据一行列的数组

	for i, colType := range colTypes {
		rowValue[i] = reflect.New(colType.ScanType())           // 跟据数据库参数类型，创建默认值 和类型
		rowParam[i] = reflect.ValueOf(&rowValue[i]).Interface() // 跟据接收的数据的类型反射出值的地址
	}
	record := make(map[string]interface{})
	// 遍历
	for rows.Next() {
		rows.Scan(rowParam...)
		for i, colType := range colTypes {
			switch colType.DatabaseTypeName() {
			case "INT", "MEDIUMINT", "SMALLINT", "TINYINT", "BIGINT":
				if rowValue[i] == nil {
					record[colType.Name()] = 0
				} else {
					record[colType.Name()] = rowValue[i]
				}
			case "DOUBLE", "FLOAT", "DECIMAL":
				if rowValue[i] == nil {
					record[colType.Name()] = 0.00
				} else {
					record[colType.Name()] = rowValue[i]
				}
			case "JSON":
				if rowValue[i] == nil {
					record[colType.Name()] = nil
				} else {
					record[colType.Name()] = rowValue[i]
				}
			default:
				if rowValue[i] == nil {
					record[colType.Name()] = ""
				} else {
					record[colType.Name()] = rowValue[i]
				}
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
func MyRowsToMap(rows *sql.Rows) []map[string]interface{} {
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
			switch colType.DatabaseTypeName() {
			case "INT", "MEDIUMINT", "SMALLINT", "TINYINT", "BIGINT":
				if rowValue[i] == nil {
					record[colType.Name()] = 0
				} else {
					record[colType.Name()] = rowValue[i]
				}
			case "DOUBLE", "FLOAT", "DECIMAL":
				if rowValue[i] == nil {
					record[colType.Name()] = 0.00
				} else {
					record[colType.Name()] = rowValue[i]
				}
			case "JSON":
				if rowValue[i] == nil {
					record[colType.Name()] = nil
				} else {
					record[colType.Name()] = rowValue[i]
				}
			default:
				if rowValue[i] == nil {
					record[colType.Name()] = ""
				} else {
					record[colType.Name()] = rowValue[i]
				}
			}
		}
		res = append(res, record)
	}
	return res
}
