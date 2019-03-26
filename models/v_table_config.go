package models

import "time"

type VTableConfig struct {
	Id_RENAME      int       `orm:"column(id)" description:"序号"`
	FieldName      string    `orm:"column(field_name);size(50);null" description:"字段"`
	FieldDesc      string    `orm:"column(field_desc);size(50);null" description:"字段名称"`
	FunctionTable1 string    `orm:"column(function_table1);size(50);null" description:"表名"`
	FunctionTable2 string    `orm:"column(function_table2);size(50);null"`
	FunctionTable3 string    `orm:"column(function_table3);size(50);null"`
	Rowver         time.Time `orm:"column(rowver);type(timestamp)" description:"时间戳"`
}
