package models

type Vcustomerformeter struct {
	CustomerId        string `orm:"column(customer_id);size(38);null"`
	CustomerName      string `orm:"column(customer_name);size(100);null" description:"客户名称"`
	DtuNo             string `orm:"column(dtu_no);size(20);null" description:"dtu编号"`
	MeterAddress      int    `orm:"column(meter_address);null" description:"表地址"`
	MeterTypeNo       string `orm:"column(meter_type_no);size(20);null" description:"电表类型编号"`
	CollectConfigName string `orm:"column(collect_config_name);size(200);null" description:"低压回路名称"`
}
