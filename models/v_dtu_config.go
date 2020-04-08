package models

import "time"

type VDtuConfig struct {
	Id_RENAME    int       `orm:"column(id)" description:"序号"`
	DtuNo        string    `orm:"column(dtu_no);size(20)" description:"dtu编号"`
	TimeInterval int       `orm:"column(time_interval);null" description:"采集间隔"`
	Rowver       time.Time `orm:"column(rowver);type(timestamp)" description:"时间戳"`
}
