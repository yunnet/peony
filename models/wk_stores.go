package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type WkStores struct {
	Id        int     `orm:"column(id);auto" description:"序号"`
	GoodName  string  `orm:"column(goodName);size(50);null" description:"商品名称"`
	Price     float64 `orm:"digits(12);decimals(2);column(price);null" description:"商品价格"`
	SpeedRate float64 `orm:"digits(12);decimals(2);column(speedrate);null" description:"速率"`
	Multiple1 int64   `orm:"column(multiple1);null" description:"商品价格倍数"`
	Multiple2 int64   `orm:"column(multiple2);null" description:"速率倍数"`
	Readonly  int     `orm:"column(readonly);null" description:"是否可用"`
}

type WkStoresAdvance struct {
	Id       int    `orm:"column(id);auto" description:"序号"`
	GoodName string `orm:"column(goodName);size(50);null" description:"商品名称"`
	Memo1    string `orm:"column(memo1);size(256)" description:"备注1"`
	Memo2    string `orm:"column(memo2);size(256)" description:"备注2"`
}

func (c *WkStores) TableName() string {
	return "wk_stores"
}

func init() {
	orm.RegisterModel(new(WkStores))
}

// AddWkStores insert a new WkStores into database and returns
// last inserted Id on success.
func AddWkStores(m *WkStores) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetWkStoresById retrieves WkStores by Id. Returns error if
// Id doesn't exist
func GetWkStoresById(id int) (v *WkStores, err error) {
	o := orm.NewOrm()
	v = &WkStores{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetStoreAdvance retrieves all advance wkStores matches certain condition. Returns empty list if
// no records exist
func GetStoreAdvance() (ml []interface{}, err error) {
	data := make([]*WkStoresAdvance, 0)
	o := orm.NewOrm()

	sqlText := "select id, goodName, memo1, memo2 from v_wk_storesbyadvance"

	if _, err = o.Raw(sqlText).QueryRows(&data); err == nil {
		for _, v := range data {
			ml = append(ml, v)
		}
		return ml, nil
	}

	return ml, nil
}

// GetAllWkStores retrieves all WkStores matches certain condition. Returns empty list if
// no records exist
func GetAllWkStores(level int) (ml []interface{}, err error) {
	data := make([]*WkStores, 0)
	o := orm.NewOrm()

	if level > 37 {
		level = 37
	}

	sqlText := "SELECT id, goodName, price, speedrate, multiple1, multiple2, readonly FROM v_wk_stores"
	if _, err = o.Raw(sqlText).QueryRows(&data); err == nil {
		for _, v := range data {
			if level < 6 {
				if v.Id == 1 {
					v.Readonly = 0
				} else {
					v.Readonly = 1
				}
			} else {
				if v.Id <= (level - 4) {
					v.Readonly = 0
				} else {
					v.Readonly = 1
				}
			}
			ml = append(ml, v)
		}
		return ml, nil
	}

	return ml, nil
}

// UpdateWkStoresById updates WkStores by Id and returns error if
// the record to be updated doesn't exist
func UpdateWkStoresById(m *WkStores) (err error) {
	o := orm.NewOrm()
	v := WkStores{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteWkStores deletes WkStores by Id and returns error if
// the record to be deleted doesn't exist
func DeleteWkStores(id int) (err error) {
	o := orm.NewOrm()
	v := WkStores{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&WkStores{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
