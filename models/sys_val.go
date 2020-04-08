package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SysVal struct {
	Id         int       `orm:"column(id);auto" description:"序号"`
	Code       string    `orm:"column(code);size(20);null" description:"代码"`
	Desc       string    `orm:"column(desc);size(60);null" description:"描述"`
	Value      string    `orm:"column(value);size(20);null" description:"变量"`
	Uplimit    string    `orm:"column(uplimit);size(20);null" description:"上限"`
	Step       string    `orm:"column(step);size(10);null" description:"步长"`
	Tag        int       `orm:"column(tag);null" description:"状态:0.正常1.删除"`
	Createuser string    `orm:"column(createuser);size(20);null" description:"创建人"`
	Createdate time.Time `orm:"column(createdate);type(datetime);null;auto_now_add" description:"创建时间"`
	Changeuser string    `orm:"column(changeuser);size(20);null" description:"修改人"`
	Changedate time.Time `orm:"column(changedate);type(datetime);null;auto_now_add" description:"修改时间"`
}

func (t *SysVal) TableName() string {
	return "sys_val"
}

func init() {
	orm.RegisterModel(new(SysVal))
}

// AddSysVal insert a new SysVal into database and returns
// last inserted Id on success.
func AddSysVal(m *SysVal) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysValById retrieves SysVal by Id. Returns error if
// Id doesn't exist
func GetSysValById(id int) (v *SysVal, err error) {
	o := orm.NewOrm()
	v = &SysVal{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysVal retrieves all SysVal matches certain condition. Returns empty list if
// no records exist
func GetAllSysVal(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysVal))
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

	var l []SysVal
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

// UpdateSysVal updates SysVal by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysValById(m *SysVal) (err error) {
	o := orm.NewOrm()
	v := SysVal{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysVal deletes SysVal by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysVal(id int) (err error) {
	o := orm.NewOrm()
	v := SysVal{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysVal{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
