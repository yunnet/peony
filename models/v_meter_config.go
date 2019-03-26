package models

import "time"

type VMeterConfig struct {
	Id_RENAME      int       `orm:"column(id)" description:"序号"`
	DtuNo          string    `orm:"column(dtu_no);size(20)" description:"dtu编号"`
	MeterAddress   int       `orm:"column(meter_address)" description:"表地址"`
	MeterTypeNo    string    `orm:"column(meter_type_no);size(20)" description:"电表类型编号"`
	GatewayNo      string    `orm:"column(gateway_no);size(20);null" description:"网关编号"`
	GatewayQzone   int       `orm:"column(gateway_qzone);null" description:"网关区域编码"`
	GatewayAddress int       `orm:"column(gateway_address);null" description:"网关地址"`
	GatewaySite    int       `orm:"column(gateway_site);null" description:"网关主站标识"`
	Pt             int       `orm:"column(pt);null" description:"电压互感器变比"`
	Ct             int       `orm:"column(ct);null" description:"电流互感器变比"`
	Rowver         time.Time `orm:"column(rowver);type(timestamp)" description:"时间戳"`
}
