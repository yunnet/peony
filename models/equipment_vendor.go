package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type EquipmentVendor struct {
	Id         int       `orm:"column(id);auto" description:"序号"`
	VendorNo   string    `orm:"column(vendor_no);size(20);null" description:"供应商编号"`
	VendorDesc string    `orm:"column(vendor_desc);size(38);null" description:"供应商名称"`
	Tag        int       `orm:"column(tag);null" description:"属性:0.正常1.删除"`
	Createuser string    `orm:"column(createuser);size(20);null" description:"创建人"`
	Createdate time.Time `orm:"column(createdate);type(datetime);null;auto_now_add" description:"创建时间"`
	Changeuser string    `orm:"column(changeuser);size(20);null" description:"修改人"`
	Changedate time.Time `orm:"column(changedate);type(datetime);null;auto_now_add" description:"修改时间"`
	Rowver     time.Time `orm:"column(rowver);type(timestamp);auto_now" description:"时间戳"`
}

func (t *EquipmentVendor) TableName() string {
	return "equipment_vendor"
}

func init() {
	orm.RegisterModel(new(EquipmentVendor))
}

// AddEquipmentVendor insert a new EquipmentVendor into database and returns
// last inserted Id on success.
func AddEquipmentVendor(m *EquipmentVendor) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEquipmentVendorById retrieves EquipmentVendor by Id. Returns error if
// Id doesn't exist
func GetEquipmentVendorById(id int) (v *EquipmentVendor, err error) {
	o := orm.NewOrm()
	v = &EquipmentVendor{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEquipmentVendor retrieves all EquipmentVendor matches certain condition. Returns empty list if
// no records exist
func GetAllEquipmentVendor(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, paging bool) (ml []interface{}, total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EquipmentVendor))
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

	var l []EquipmentVendor
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
		return ml, total, nil
	}
	return nil, 0, err
}

// UpdateEquipmentVendor updates EquipmentVendor by Id and returns error if
// the record to be updated doesn't exist
func UpdateEquipmentVendorById(m *EquipmentVendor) (err error) {
	o := orm.NewOrm()
	v := EquipmentVendor{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEquipmentVendor deletes EquipmentVendor by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEquipmentVendor(id int) (err error) {
	o := orm.NewOrm()
	v := EquipmentVendor{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EquipmentVendor{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
