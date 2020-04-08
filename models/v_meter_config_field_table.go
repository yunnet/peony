package models

import "time"

type VMeterConfigFieldTable struct {
	Id_RENAME      int       `orm:"column(id)" description:"序号"`
	DtuNo          string    `orm:"column(dtu_no);size(20)" description:"dtu编号"`
	MeterAddress   int       `orm:"column(meter_address)" description:"表地址"`
	DsAddr         int       `orm:"column(ds_addr);null" description:"网关传输地址"`
	FieldName      string    `orm:"column(field_name);size(50);null" description:"字段"`
	FunctionTable1 string    `orm:"column(function_table1);size(50);null" description:"表名"`
	FunctionTable2 string    `orm:"column(function_table2);size(50);null"`
	FunctionTable3 string    `orm:"column(function_table3);size(50);null"`
	Rowver         time.Time `orm:"column(rowver);type(timestamp)" description:"时间戳"`
}
