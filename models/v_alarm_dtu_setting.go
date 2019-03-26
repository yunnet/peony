package models

import "time"

type VAlarmDtuSetting struct {
	Id_RENAME    int       `orm:"column(id)" description:"序号"`
	DtuNo        string    `orm:"column(dtu_no);size(20)" description:"dtu编号"`
	MeterAddress int       `orm:"column(meter_address)" description:"表地址"`
	AlarmTypeId  string    `orm:"column(alarm_type_id);size(20)" description:"报警类型编号"`
	MinValue     float64   `orm:"column(min_value);null;digits(20);decimals(4)" description:"最小值"`
	MaxValue     float64   `orm:"column(max_value);null;digits(20);decimals(4)" description:"最大值"`
	Rowver       time.Time `orm:"column(rowver);type(timestamp)" description:"时间戳"`
}
