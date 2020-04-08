package models

type VFuncTools struct {
	MeterTypeNo    string  `orm:"column(meter_type_no);size(20);null" description:"电表类型编号"`
	AddressSort    int     `orm:"column(address_sort);null" description:"地址排序"`
	RomAddress     string  `orm:"column(rom_address);size(40);null" description:"寄存器地址"`
	RomName        string  `orm:"column(rom_name);size(80);null" description:"寄存器名称"`
	Units          string  `orm:"column(units);size(40);null" description:"单位"`
	DataType       string  `orm:"column(data_type);size(40);null" description:"数据类型"`
	DataValue      float64 `orm:"column(data_value);digits(10);decimals(6)"`
	Calcfactor     float64 `orm:"column(calcfactor);null;digits(10);decimals(6)" description:"计算因子"`
	ByteLength     int     `orm:"column(byte_length);null" description:"数据长度"`
	NeedPt         int8    `orm:"column(need_pt);null" description:"是否电压互感器变化"`
	NeedCt         int8    `orm:"column(need_ct);null" description:"是否电流互感器变比"`
	SegmentNo      int     `orm:"column(segment_no);null" description:"表示该属性属于协议的第几段"`
	OFFSET         int     `orm:"column(OFFSET);null" description:"解析协议时相对上个属性的偏移值"`
	FunctionTable1 string  `orm:"column(function_table1);size(50);null" description:"基础表"`
	FunctionTable2 string  `orm:"column(function_table2);size(50);null" description:"附加表"`
	FunctionTable3 string  `orm:"column(function_table3);size(50);null" description:"业务表"`
	FunctionField  string  `orm:"column(function_field);size(50);null" description:"字段名"`
}
