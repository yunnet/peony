package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SysConnConfig struct {
	Id               int       `orm:"column(PK_ID);pk"`
	DBTYPE           string    `orm:"column(DB_TYPE);size(50);null"`
	HOST             string    `orm:"column(HOST);size(50);null"`
	PORT             int       `orm:"column(PORT);null"`
	USERNAME         string    `orm:"column(USER_NAME);size(30);null"`
	PASSWORD         string    `orm:"column(PASSWORD);size(256);null"`
	SHORTNAME        string    `orm:"column(SHORT_NAME);size(100);null"`
	REMARK           string    `orm:"column(REMARK);size(1024);null"`
	DBNAME           string    `orm:"column(DB_NAME);size(50);null"`
	CREATEPERSONNAME string    `orm:"column(CREATE_PERSON_NAME);size(50);null"`
	CREATETIME       time.Time `orm:"column(CREATE_TIME);type(datetime);null"`
	DELETED          int       `orm:"column(DELETED);null"`
	DBVERSION        string    `orm:"column(DB_VERSION);size(30);null"`
}

func (t *SysConnConfig) TableName() string {
	return "sys_conn_config"
}

func init() {
	orm.RegisterModel(new(SysConnConfig))
}

// AddSysConnConfig insert a new SysConnConfig into database and returns
// last inserted Id on success.
func AddSysConnConfig(m *SysConnConfig) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysConnConfigById retrieves SysConnConfig by Id. Returns error if
// Id doesn't exist
func GetSysConnConfigById(id int) (v *SysConnConfig, err error) {
	o := orm.NewOrm()
	v = &SysConnConfig{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysConnConfig retrieves all SysConnConfig matches certain condition. Returns empty list if
// no records exist
func GetAllSysConnConfig(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysConnConfig))
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

	var l []SysConnConfig
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

// UpdateSysConnConfig updates SysConnConfig by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysConnConfigById(m *SysConnConfig) (err error) {
	o := orm.NewOrm()
	v := SysConnConfig{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysConnConfig deletes SysConnConfig by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysConnConfig(id int) (err error) {
	o := orm.NewOrm()
	v := SysConnConfig{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysConnConfig{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
