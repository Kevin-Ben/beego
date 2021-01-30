package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Content struct {
	Id        int64     `orm:"column(id);pk" description:"主健"`
	SubPri    string    `orm:"column(sub_pri);size(100)" description:"订阅"`
	Title     string    `orm:"column(title);size(100)" description:"标题"`
	SendTo    string    `orm:"column(send_to);size(100)" description:"接收人"`
	Content   string    `orm:"column(content)" description:"内容"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime)" description:"创建时间"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime)" description:"更新时间"`
}

func (t *Content) TableName() string {
	return "content"
}

func init() {
	orm.RegisterModelWithPrefix("ue_", new(Content))
}

// AddContent insert a new Content into database and returns
// last inserted Id on success.
func AddContent(m *Content) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return id, err
}

// GetContentById retrieves Content by Id. Returns error if
// Id doesn't exist
func GetContentById(id int64) (v *Content, err error) {
	o := orm.NewOrm()
	v = &Content{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllContent retrieves all Content matches certain condition. Returns empty list if
// no records exist
func GetAllContent(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Content))
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

	var l []Content
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

// UpdateContent updates Content by Id and returns error if
// the record to be updated doesn't exist
func UpdateContentById(m *Content) (err error) {
	o := orm.NewOrm()
	v := Content{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteContent deletes Content by Id and returns error if
// the record to be deleted doesn't exist
func DeleteContent(id int64) (err error) {
	o := orm.NewOrm()
	v := Content{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Content{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
