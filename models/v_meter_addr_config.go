package models

import "time"

type VMeterAddrConfig struct {
	Id_RENAME        int       `orm:"column(id)" description:"序号"`
	MeterTypeNo      string    `orm:"column(meter_type_no);size(20)" description:"电表类型编号"`
	SegmentStartAddr int       `orm:"column(segment_start_addr);null" description:"分段起始地址"`
	SegmentLen       int       `orm:"column(segment_len);null" description:"读取寄存器长度"`
	SegmentNo        int       `orm:"column(segment_no);null" description:"第几段"`
	Rowver           time.Time `orm:"column(rowver);type(timestamp)" description:"时间戳"`
}
