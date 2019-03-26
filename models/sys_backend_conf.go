package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SysBackendConf struct {
	Id         int       `orm:"column(id);auto" description:"序号"`
	Appname    string    `orm:"column(appname);size(50)"`
	Appversion string    `orm:"column(appversion);size(50)"`
	Deploy     string    `orm:"column(deploy);size(50)"`
	Configtext string    `orm:"column(configtext);size(5000);null"`
	Tag        int       `orm:"column(tag);null" description:"状态"`
	Createuser string    `orm:"column(createuser);size(20);null" description:"创建人"`
	Createdate time.Time `orm:"column(createdate);type(datetime);null;auto_now_add" description:"创建时间"`
	Changeuser string    `orm:"column(changeuser);size(20);null" description:"修改人"`
	Changedate time.Time `orm:"column(changedate);type(datetime);null;auto_now_add" description:"修改时间"`
}

func (t *SysBackendConf) TableName() string {
	return "sys_backend_conf"
}

func init() {
	orm.RegisterModel(new(SysBackendConf))
}

// AddSysBackendConf insert a new SysBackendConf into database and returns
// last inserted Id on success.
func AddSysBackendConf(m *SysBackendConf) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysBackendConfById retrieves SysBackendConf by Id. Returns error if
// Id doesn't exist
func GetSysBackendConfById(id int) (v *SysBackendConf, err error) {
	o := orm.NewOrm()
	v = &SysBackendConf{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysBackendConf retrieves all SysBackendConf matches certain condition. Returns empty list if
// no records exist
func GetAllSysBackendConf(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysBackendConf))
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

	var l []SysBackendConf
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

// UpdateSysBackendConf updates SysBackendConf by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysBackendConfById(m *SysBackendConf) (err error) {
	o := orm.NewOrm()
	v := SysBackendConf{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysBackendConf deletes SysBackendConf by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysBackendConf(id int) (err error) {
	o := orm.NewOrm()
	v := SysBackendConf{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysBackendConf{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
