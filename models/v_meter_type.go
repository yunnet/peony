package models

import "time"

type VMeterType struct {
	Id_RENAME   int       `orm:"column(id)" description:"序号"`
	MeterTypeNo string    `orm:"column(meter_type_no);size(20);null" description:"电表类型编号"`
	MeterType   string    `orm:"column(meter_type);size(40);null" description:"电表类型"`
	PtAddress   string    `orm:"column(pt_address);size(20);null" description:"电压互感器变比地址"`
	CtAddress   string    `orm:"column(ct_address);size(20);null" description:"电流互感器变比地址"`
	ThreePhase  int8      `orm:"column(three_phase);null" description:"是否三相电表:0.单相1.三相"`
	Rowver      time.Time `orm:"column(rowver);type(timestamp)" description:"时间戳"`
}
