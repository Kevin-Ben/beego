package models

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/golang-module/carbon"
	"reflect"
	"strings"
	"time"
)

var (
	Subscribes map[string]*Subscribe
)

type Subscribe struct {
	Id        int64     `orm:"auto" type:"pk"`
	SubPri    string    `orm:"size(100)"`
	SendTo    string    `orm:"size(100)"`
	CreatedAt time.Time `orm:"type(datetime)"`
	UpdatedAt time.Time `orm:"type(datetime)"`
}

func init() {
	orm.RegisterModelWithPrefix("ue_", new(Subscribe))
}

func ReadOrCreateSubscribe(subPri string, sendTo string) (v *Subscribe, err error) {
	var objecId int64
	o := orm.NewOrm()
	object := &Subscribe{SubPri: subPri, SendTo: sendTo}
	if err := o.Read(object); err == nil {
		return v, nil
	}

	object.CreatedAt = carbon.Now().ToGoTime()
	object.UpdatedAt = object.CreatedAt
	objecId, err = o.Insert(object)
	object.Id = int64(objecId)
	return object, err

}

// AddSubscribe insert a new Subscribe into database and returns
// last inserted Id on success.
func AddSubscribe(m *Subscribe) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetSubscribeBySubPri(search string) (v *Subscribe, err error) {
	o := orm.NewOrm()
	v = &Subscribe{SubPri: search}
	if err = o.QueryTable(new(Subscribe)).Filter("sub_pri", search).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetSubscribeById retrieves Subscribe by Id. Returns error if
// Id doesn't exist
func GetSubscribeById(id int64) (v *Subscribe, err error) {
	o := orm.NewOrm()
	v = &Subscribe{Id: id}
	if err = o.QueryTable(new(Subscribe)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSubscribe retrieves all Subscribe matches certain condition. Returns empty list if
// no records exist
func GetAllSubscribe(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Subscribe))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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

	var l []Subscribe
	qs = qs.OrderBy(sortFields...).RelatedSel()
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

// UpdateSubscribe updates Subscribe by Id and returns error if
// the record to be updated doesn't exist
func UpdateSubscribeById(m *Subscribe) (err error) {
	o := orm.NewOrm()
	v := Subscribe{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSubscribe deletes Subscribe by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSubscribe(id int64) (err error) {
	o := orm.NewOrm()
	v := Subscribe{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Subscribe{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
