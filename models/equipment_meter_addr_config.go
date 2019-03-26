package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type EquipmentMeterAddrConfig struct {
	Id               int       `orm:"column(id);auto" description:"序号"`
	MeterTypeNo      string    `orm:"column(meter_type_no);size(20)" description:"电表类型编号"`
	SegmentStartAddr int       `orm:"column(segment_start_addr);null" description:"分段起始地址"`
	SegmentLen       int       `orm:"column(segment_len);null" description:"读取寄存器长度"`
	SegmentNo        int       `orm:"column(segment_no);null" description:"第几段"`
	Tag              int       `orm:"column(tag);null" description:"状态:0.正常1.删除"`
	Createuser       string    `orm:"column(createuser);size(20);null" description:"创建人"`
	Createdate       time.Time `orm:"column(createdate);type(datetime);null;auto_now_add" description:"创建时间"`
	Changeuser       string    `orm:"column(changeuser);size(20);null" description:"修改人"`
	Changedate       time.Time `orm:"column(changedate);type(datetime);null;auto_now_add" description:"修改时间"`
	Rowver           time.Time `orm:"column(rowver);type(timestamp);auto_now" description:"时间戳"`
}

func (t *EquipmentMeterAddrConfig) TableName() string {
	return "equipment_meter_addr_config"
}

func init() {
	orm.RegisterModel(new(EquipmentMeterAddrConfig))
}

// AddEquipmentMeterAddrConfig insert a new EquipmentMeterAddrConfig into database and returns
// last inserted Id on success.
func AddEquipmentMeterAddrConfig(m *EquipmentMeterAddrConfig) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEquipmentMeterAddrConfigById retrieves EquipmentMeterAddrConfig by Id. Returns error if
// Id doesn't exist
func GetEquipmentMeterAddrConfigById(id int) (v *EquipmentMeterAddrConfig, err error) {
	o := orm.NewOrm()
	v = &EquipmentMeterAddrConfig{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEquipmentMeterAddrConfig retrieves all EquipmentMeterAddrConfig matches certain condition. Returns empty list if
// no records exist
func GetAllEquipmentMeterAddrConfig(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EquipmentMeterAddrConfig))
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []EquipmentMeterAddrConfig
	qs = qs.OrderBy(sortFields...)
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
		return ml, nil
	}
	return nil, err
}

// UpdateEquipmentMeterAddrConfig updates EquipmentMeterAddrConfig by Id and returns error if
// the record to be updated doesn't exist
func UpdateEquipmentMeterAddrConfigById(m *EquipmentMeterAddrConfig) (err error) {
	o := orm.NewOrm()
	v := EquipmentMeterAddrConfig{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEquipmentMeterAddrConfig deletes EquipmentMeterAddrConfig by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEquipmentMeterAddrConfig(id int) (err error) {
	o := orm.NewOrm()
	v := EquipmentMeterAddrConfig{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EquipmentMeterAddrConfig{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
