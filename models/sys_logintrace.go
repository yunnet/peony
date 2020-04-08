package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SysLogintrace struct {
	Id         int       `orm:"column(id);auto" description:"序号"`
	User       string    `orm:"column(user);size(30);null" description:"用户名"`
	RemoteAddr string    `orm:"column(remoteAddr);size(50);null" description:"IP地址"`
	LoginTime  time.Time `orm:"column(loginTime);type(datetime);null" description:"登录时间"`
}

func (t *SysLogintrace) TableName() string {
	return "sys_logintrace"
}

func init() {
	orm.RegisterModel(new(SysLogintrace))
}

// AddSysLogintrace insert a new SysLogintrace into database and returns
// last inserted Id on success.
func AddSysLogintrace(m *SysLogintrace) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysLogintraceById retrieves SysLogintrace by Id. Returns error if
// Id doesn't exist
func GetSysLogintraceById(id int) (v *SysLogintrace, err error) {
	o := orm.NewOrm()
	v = &SysLogintrace{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysLogintrace retrieves all SysLogintrace matches certain condition. Returns empty list if
// no records exist
func GetAllSysLogintrace(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysLogintrace))
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

	var l []SysLogintrace
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

// UpdateSysLogintrace updates SysLogintrace by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysLogintraceById(m *SysLogintrace) (err error) {
	o := orm.NewOrm()
	v := SysLogintrace{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysLogintrace deletes SysLogintrace by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysLogintrace(id int) (err error) {
	o := orm.NewOrm()
	v := SysLogintrace{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysLogintrace{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
