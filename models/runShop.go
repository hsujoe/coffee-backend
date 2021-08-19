package models

import (
	"coffee_backend/db"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

type RunShop struct {
	Id         int    `json:"id" form:"id"`
	Store_id   int    `json:"sid" form:"sid"`
	SName      string `json:"sname" form:"sname"`
	Product_id int    `json:"pid" form:"pid"`
	PName      string `json:"pname" form:"pname"`
	Tasting    string `json:"tasting" form:"tasting"`
	Memo       string `json:"memo" form:"memo"`
	State      int    `json:"state" form:"state"`
}

//查询所有记录
func (runShop *RunShop) GetRunShops(name string) (RunShops []RunShop, err error) {
	m := map[string]interface{}{"name": name}
	var where []string
	for _, k := range []string{"name"} {
		if v, ok := m[k]; ok {
			if len(fmt.Sprintf("%v", v)) > 0 {
				where = append(where, fmt.Sprintf("%s like 'vb%svb'", k, fmt.Sprintf("%v", v)))
			}
		}
	}
	sql := "SELECT rs.id as id , s.id as sid,s.name as sname, p.id as pid,p.name as pname,rs.tasting,rs.memo as memo,rs.state as state FROM RUNSHOP rs " +
		" LEFT JOIN STORE s on s.id = rs.store_id " +
		" LEFT JOIN PRODUCT p on p.id =rs.product_id "

	if len(where) > 0 {
		sql += " where " + strings.Join(where, " AND ")
		sql = strings.ReplaceAll(sql, "vb", "%")
	}

	fmt.Print(sql)
	// rows, err := db.SqlDB.Query("select id,product_no,name from PRODUCTS where " + strings.Join(where, " AND "))
	rows, err := db.SqlDB.Query(sql)
	if err == nil {
		db.SqlDB.Begin()
		for rows.Next() {
			runShop := RunShop{}
			err := rows.Scan(&runShop.Id, &runShop.Store_id, &runShop.SName, &runShop.Product_id,
				&runShop.PName, &runShop.Tasting, &runShop.Memo, &runShop.State)
			if err != nil {
				log.Fatal(err)
			}
			RunShops = append(RunShops, runShop)
		}

		rows.Close()
	}
	return
}

//查询一条记录5
func (rs *RunShop) GetRunShop() (runShop RunShop, err error) {
	runShop = RunShop{}
	err = db.SqlDB.QueryRow("SELECT rs.id as id , s.id as sid,s.name as sname, p.id as pid,p.name as pname,rs.tasting,rs.memo as memo,rs.state as state FROM RUNSHOP rs "+
		"LEFT JOIN STORE s on s.id = rs.store_id  LEFT JOIN PRODUCT p on p.id =rs.product_id where rs.id = ?", rs.Id).
		Scan(&runShop.Id, &runShop.Store_id, &runShop.SName, &runShop.Product_id,
			&runShop.PName, &runShop.Tasting, &runShop.Memo, &runShop.State)
	return
}

//插入
func (runShop RunShop) CreateRunShop() int64 {
	t := reflect.TypeOf(runShop)
	va := reflect.ValueOf(runShop)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = va.Field(i).Interface()
	}
	var fieldData []string
	var valesData []string
	for k, _ := range data {
		if k != "SName" && k != "PName" {
			fieldData = append(fieldData, k)
		}
	}

	for _, k := range fieldData {
		if v, ok := data[k]; ok {
			if len(fmt.Sprintf("\"%v\"", v)) > 0 {
				if k == "Id" || k == "State" || k == "Sid" || k == "Pid" {
					valesData = append(valesData, fmt.Sprintf("%s", fmt.Sprintf("%v", v)))
				} else {
					if k != "SName" && k != "PName" {
						valesData = append(valesData, fmt.Sprintf("\"%v\"", v))
					}
				}

			}
		}
	}

	sql := "insert into RUNSHOP ( " + strings.Join(fieldData, ",") + " ) values ( " + strings.Join(valesData, ",") + " )"
	fmt.Print(sql)
	// rs, err := db.SqlDB.Exec("INSERT into PRODUCTS (id,product_no, name,Contrast_no) values (?,?,?,?)", product.Id, product.Product_no, product.Name, product.Contrast_no)
	rs, err := db.SqlDB.Exec(sql)

	// rs, err := db.SqlDB.ExecContext(context.Background(), sql, time.Now())

	if err != nil {
		log.Fatal(err)
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return id
}

//修改
func (runShop *RunShop) UpdateRunShop() int64 {
	mainMap := structs.Map(runShop)
	var mainValesData []string
	for key, value := range mainMap {
		if len(fmt.Sprintf("%v", value)) > 0 {
			if strings.Compare("Id", fmt.Sprintf("%v", key)) != 0 {
				mainValesData = append(mainValesData, fmt.Sprintf("%s = \"%v\" ", key, fmt.Sprintf("%v", value)))
			}
		}
	}

	tx, _ := db.SqlDB.Begin()
	_, err := db.SqlDB.Exec("update RUNSHOP set "+strings.Join(mainValesData, " , ")+" where id = ?", runShop.Id)
	if err != nil {
		defer tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
	return 0

}

//删除一条记录
func DeleteRunShop(id int) int64 {
	rs, err := db.SqlDB.Exec("delete from BREW where id = ?", id)
	if err != nil {
		log.Fatal()
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		log.Fatal()
	}
	return rows
}
