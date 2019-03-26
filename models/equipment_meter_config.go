package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type EquipmentMeterConfig struct {
	Id             int       `orm:"column(id);auto" description:"序号"`
	DtuNo          string    `orm:"column(dtu_no);size(20)" description:"dtu编号"`
	MeterAddress   int       `orm:"column(meter_address)" description:"表地址"`
	MeterTypeNo    string    `orm:"column(meter_type_no);size(20)" description:"电表类型编号"`
	GatewayNo      string    `orm:"column(gateway_no);size(20);null" description:"网关编号"`
	GatewayQzone   int       `orm:"column(gateway_qzone);null" description:"网关区域编码"`
	GatewayAddress int       `orm:"column(gateway_address);null" description:"网关地址"`
	GatewaySite    int       `orm:"column(gateway_site);null" description:"网关主站标识"`
	LoopName       string    `orm:"column(loop_name);size(30);null" description:"回路名称"`
	Pt             int       `orm:"column(pt);null" description:"电压互感器变比"`
	Ct             int       `orm:"column(ct);null" description:"电流互感器变比"`
	Tag            int       `orm:"column(tag);null" description:"状态:0.正常1.删除"`
	Createuser     string    `orm:"column(createuser);size(20);null" description:"创建人"`
	Createdate     time.Time `orm:"column(createdate);type(datetime);null;auto_now_add" description:"创建时间"`
	Changeuser     string    `orm:"column(changeuser);size(20);null" description:"修改人"`
	Changedate     time.Time `orm:"column(changedate);type(datetime);null;auto_now_add" description:"修改时间"`
	Rowver         time.Time `orm:"column(rowver);type(timestamp);auto_now" description:"时间戳"`
	CustomerId     string    `orm:"column(customer_id);size(38);null" description:"客户id"`
	Used           int       `orm:"column(used);null" description:"是否使用"`
}

func (t *EquipmentMeterConfig) TableName() string {
	return "equipment_meter_config"
}

func init() {
	orm.RegisterModel(new(EquipmentMeterConfig))
}

// AddEquipmentMeterConfig insert a new EquipmentMeterConfig into database and returns
// last inserted Id on success.
func AddEquipmentMeterConfig(m *EquipmentMeterConfig) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEquipmentMeterConfigById retrieves EquipmentMeterConfig by Id. Returns error if
// Id doesn't exist
func GetEquipmentMeterConfigById(id int) (v *EquipmentMeterConfig, err error) {
	o := orm.NewOrm()
	v = &EquipmentMeterConfig{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEquipmentMeterConfig retrieves all EquipmentMeterConfig matches certain condition. Returns empty list if
// no records exist
func GetAllEquipmentMeterConfig(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, paging bool) (ml []interface{}, total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EquipmentMeterConfig))
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

	var l []EquipmentMeterConfig
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

// UpdateEquipmentMeterConfig updates EquipmentMeterConfig by Id and returns error if
// the record to be updated doesn't exist
func UpdateEquipmentMeterConfigById(m *EquipmentMeterConfig) (err error) {
	o := orm.NewOrm()
	v := EquipmentMeterConfig{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEquipmentMeterConfig deletes EquipmentMeterConfig by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEquipmentMeterConfig(id int) (err error) {
	o := orm.NewOrm()
	v := EquipmentMeterConfig{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EquipmentMeterConfig{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
