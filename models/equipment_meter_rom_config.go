package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type EquipmentMeterRomConfig struct {
	Id             int       `orm:"column(id);auto" description:"序号"`
	MeterTypeNo    string    `orm:"column(meter_type_no);size(20);null" description:"电表类型编号"`
	AddressSort    int       `orm:"column(address_sort);null" description:"地址排序"`
	RomAddress     string    `orm:"column(rom_address);size(40);null" description:"寄存器地址"`
	RomName        string    `orm:"column(rom_name);size(80);null" description:"寄存器名称"`
	Units          string    `orm:"column(units);size(40);null" description:"单位"`
	DataType       string    `orm:"column(data_type);size(40);null" description:"数据类型"`
	SegmentNo      int       `orm:"column(segment_no);null" description:"表示该属性属于协议的第几段"`
	Offset         int       `orm:"column(offset);null" description:"解析协议时相对上个属性的偏移值"`
	NeedPt         int8      `orm:"column(need_pt);null" description:"是否电压互感器变化"`
	NeedCt         int8      `orm:"column(need_ct);null" description:"是否电流互感器变比"`
	Calcfactor     float64   `orm:"column(calcfactor);null;digits(10);decimals(6)" description:"计算因子"`
	MsbBit         int8      `orm:"column(msb_bit);null" description:"最高位为符号位"`
	ByteLength     int       `orm:"column(byte_length);null" description:"数据长度"`
	FunctionTable1 string    `orm:"column(function_table1);size(50);null" description:"基础表"`
	FunctionTable2 string    `orm:"column(function_table2);size(50);null" description:"附加表"`
	FunctionTable3 string    `orm:"column(function_table3);size(50);null" description:"业务表"`
	FunctionField  string    `orm:"column(function_field);size(50);null" description:"字段名"`
	Tag            int       `orm:"column(tag);null" description:"状态:0.正常1.删除"`
	Createuser     string    `orm:"column(createuser);size(20);null" description:"创建人"`
	Createdate     time.Time `orm:"column(createdate);type(datetime);null;auto_now_add" description:"创建时间"`
	Changeuser     string    `orm:"column(changeuser);size(20);null" description:"修改人"`
	Changedate     time.Time `orm:"column(changedate);type(datetime);null;auto_now_add" description:"修改时间"`
	Rowver         time.Time `orm:"column(rowver);type(timestamp);auto_now" description:"时间戳"`
}

func (t *EquipmentMeterRomConfig) TableName() string {
	return "equipment_meter_rom_config"
}

func init() {
	orm.RegisterModel(new(EquipmentMeterRomConfig))
}

// AddEquipmentMeterRomConfig insert a new EquipmentMeterRomConfig into database and returns
// last inserted Id on success.
func AddEquipmentMeterRomConfig(m *EquipmentMeterRomConfig) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEquipmentMeterRomConfigById retrieves EquipmentMeterRomConfig by Id. Returns error if
// Id doesn't exist
func GetEquipmentMeterRomConfigById(id int) (v *EquipmentMeterRomConfig, err error) {
	o := orm.NewOrm()
	v = &EquipmentMeterRomConfig{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEquipmentMeterRomConfig retrieves all EquipmentMeterRomConfig matches certain condition. Returns empty list if
// no records exist
func GetAllEquipmentMeterRomConfig(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, paging bool) (ml []interface{}, total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EquipmentMeterRomConfig))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, 0, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, 0, errors.New("Error: unused 'order' fields")
		}
	}

	var l []EquipmentMeterRomConfig
	qs = qs.OrderBy(sortFields...)
	if paging{
		total, _ = qs.Count()
	}

	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, total,nil
	}
	return nil, 0, err
}

// UpdateEquipmentMeterRomConfig updates EquipmentMeterRomConfig by Id and returns error if
// the record to be updated doesn't exist
func UpdateEquipmentMeterRomConfigById(m *EquipmentMeterRomConfig) (err error) {
	o := orm.NewOrm()
	v := EquipmentMeterRomConfig{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEquipmentMeterRomConfig deletes EquipmentMeterRomConfig by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEquipmentMeterRomConfig(id int) (err error) {
	o := orm.NewOrm()
	v := EquipmentMeterRomConfig{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EquipmentMeterRomConfig{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
