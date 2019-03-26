package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SysBackendUser struct {
	Id       int    `orm:"column(id);auto"`
	RealName string `orm:"column(real_name);size(255)"`
	UserName string `orm:"column(user_name);size(255)"`
	UserPwd  string `orm:"column(user_pwd);size(255)"`
	IsSuper  int8   `orm:"column(is_super)"`
	Status   int    `orm:"column(status)"`
	Mobile   string `orm:"column(mobile);size(16)"`
	Email    string `orm:"column(email);size(256)"`
	Avatar   string `orm:"column(avatar);size(256)"`

}

func (t *SysBackendUser) TableName() string {
	return "sys_backend_user"
}

func init() {
	orm.RegisterModel(new(SysBackendUser))
}

// AddSysBackendUser insert a new SysBackendUser into database and returns
// last inserted Id on success.
func AddSysBackendUser(m *SysBackendUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysBackendUserById retrieves SysBackendUser by Id. Returns error if
// Id doesn't exist
func GetSysBackendUserById(id int) (v *SysBackendUser, err error) {
	o := orm.NewOrm()
	v = &SysBackendUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetSysBackendUserByName(username, password string)(v *SysBackendUser, err error) {
	o := orm.NewOrm()
	v = &SysBackendUser{UserName: username, UserPwd: password}
	if err = o.Read(v, "user_name", "user_pwd"); err == nil{
		return v, nil
	}
	return nil, err
}

// GetAllSysBackendUser retrieves all SysBackendUser matches certain condition. Returns empty list if
// no records exist
func GetAllSysBackendUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysBackendUser))
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

	var l []SysBackendUser
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

// UpdateSysBackendUser updates SysBackendUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysBackendUserById(m *SysBackendUser) (err error) {
	o := orm.NewOrm()
	v := SysBackendUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysBackendUser deletes SysBackendUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysBackendUser(id int) (err error) {
	o := orm.NewOrm()
	v := SysBackendUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysBackendUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
