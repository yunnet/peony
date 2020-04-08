package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type EquipmentTableConfig struct {
	Id             int       `orm:"column(id);auto" description:"序号"`
	FieldName      string    `orm:"column(field_name);size(50);null" description:"字段"`
	FieldDesc      string    `orm:"column(field_desc);size(50);null" description:"字段名称"`
	FunctionTable1 string    `orm:"column(function_table1);size(50);null" description:"表名"`
	FunctionTable2 string    `orm:"column(function_table2);size(50);null"`
	FunctionTable3 string    `orm:"column(function_table3);size(50);null"`
	Tag            int       `orm:"column(tag);null" description:"状态:0.正常1.删除"`
	Createuser     string    `orm:"column(createuser);size(20);null" description:"创建人"`
	Createdate     time.Time `orm:"column(createdate);type(datetime);null;auto_now_add" description:"创建时间"`
	Changeuser     string    `orm:"column(changeuser);size(20);null" description:"修改人"`
	Changedate     time.Time `orm:"column(changedate);type(datetime);null;auto_now_add" description:"修改时间"`
	Rowver         time.Time `orm:"column(rowver);type(timestamp);auto_now" description:"时间戳"`
}

func (t *EquipmentTableConfig) TableName() string {
	return "equipment_table_config"
}

func init() {
	orm.RegisterModel(new(EquipmentTableConfig))
}

// AddEquipmentTableConfig insert a new EquipmentTableConfig into database and returns
// last inserted Id on success.
func AddEquipmentTableConfig(m *EquipmentTableConfig) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEquipmentTableConfigById retrieves EquipmentTableConfig by Id. Returns error if
// Id doesn't exist
func GetEquipmentTableConfigById(id int) (v *EquipmentTableConfig, err error) {
	o := orm.NewOrm()
	v = &EquipmentTableConfig{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEquipmentTableConfig retrieves all EquipmentTableConfig matches certain condition. Returns empty list if
// no records exist
func GetAllEquipmentTableConfig(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EquipmentTableConfig))
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

	var l []EquipmentTableConfig
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

// UpdateEquipmentTableConfig updates EquipmentTableConfig by Id and returns error if
// the record to be updated doesn't exist
func UpdateEquipmentTableConfigById(m *EquipmentTableConfig) (err error) {
	o := orm.NewOrm()
	v := EquipmentTableConfig{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEquipmentTableConfig deletes EquipmentTableConfig by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEquipmentTableConfig(id int) (err error) {
	o := orm.NewOrm()
	v := EquipmentTableConfig{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EquipmentTableConfig{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
