package models

import (
	"coffee_backend/db"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

type Product struct {
	Id    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Unit  string `json:"unit" form:"unit"`
	Price int    `json:"price" form:"price"`
	Image string `json:"image" form:"image"`
	Memo  string `json:"memo" form:"memo"`
	State int    `json:"state" form:"state"`
}

//查询所有记录
func (product *Product) GetRows(name string) (products []Product, err error) {
	m := map[string]interface{}{"name": name}
	var where []string
	for _, k := range []string{"name"} {
		if v, ok := m[k]; ok {
			if len(fmt.Sprintf("%v", v)) > 0 {
				where = append(where, fmt.Sprintf("%s like 'vb%svb'", k, fmt.Sprintf("%v", v)))
			}
		}
	}
	sql := "select id,name,unit,price,memo,state,image from PRODUCT where state in (0,1)"
	if len(where) > 0 {
		sql += strings.Join(where, " AND ")
		sql = strings.ReplaceAll(sql, "vb", "%")
	}

	fmt.Print(sql)
	// rows, err := db.SqlDB.Query("select id,product_no,name from PRODUCTS where " + strings.Join(where, " AND "))
	rows, err := db.SqlDB.Query(sql)
	db.SqlDB.Begin()
	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Unit, &product.Price, &product.Memo, &product.State, &product.Image)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	rows.Close()
	return
}

//查询推薦
func (product *Product) GetRecommendRows() (products []Product, err error) {
	sql := "select id,name,image from PRODUCT where state =2"

	fmt.Print(sql)
	// rows, err := db.SqlDB.Query("select id,product_no,name from PRODUCTS where " + strings.Join(where, " AND "))
	rows, err := db.SqlDB.Query(sql)
	db.SqlDB.Begin()
	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Image)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	rows.Close()
	return
}

//查询一条记录5
func (p *Product) GetRow() (product Product, err error) {
	product = Product{}
	err = db.SqlDB.QueryRow("select id,name,unit,price,memo,state,image from PRODUCT where id = ?", p.Id).
		Scan(&product.Id, &product.Name, &product.Unit, &product.Price, &product.Memo, &product.State, &product.Image)
	return
}

//插入
func (product Product) Create() int64 {
	t := reflect.TypeOf(product)
	va := reflect.ValueOf(product)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = va.Field(i).Interface()
	}
	var fieldData []string
	var valesData []string
	for k, _ := range data {
		fieldData = append(fieldData, k)
	}

	for _, k := range fieldData {
		if v, ok := data[k]; ok {
			if len(fmt.Sprintf("\"%v\"", v)) > 0 {
				if k == "Id" || k == "Price" || k == "State" {
					valesData = append(valesData, fmt.Sprintf("%s", fmt.Sprintf("%v", v)))
				} else {
					valesData = append(valesData, fmt.Sprintf("\"%v\"", v))
				}

			}
		}
	}

	sql := "insert into PRODUCT ( " + strings.Join(fieldData, ",") + " ) values ( " + strings.Join(valesData, ",") + " )"
	fmt.Print(sql)
	// rs, err := db.SqlDB.Exec("INSERT into PRODUCTS (id,product_no, name,Contrast_no) values (?,?,?,?)", product.Id, product.Product_no, product.Name, product.Contrast_no)
	rs, err := db.SqlDB.Exec(sql)
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
func (product *Product) Update() int64 {
	mainMap := structs.Map(product)
	var mainValesData []string
	for key, value := range mainMap {
		if len(fmt.Sprintf("%v", value)) > 0 {
			if strings.Compare("Id", fmt.Sprintf("%v", key)) != 0 {
				mainValesData = append(mainValesData, fmt.Sprintf("%s = \"%v\" ", key, fmt.Sprintf("%v", value)))
			}
		}
	}

	tx, _ := db.SqlDB.Begin()
	_, err := db.SqlDB.Exec("update PRODUCT set "+strings.Join(mainValesData, " , ")+" where id = ?", product.Id)
	if err != nil {
		defer tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
	return 0

}

//删除一条记录
func Delete(id int) int64 {
	rs, err := db.SqlDB.Exec("delete from PRODUCT where id = ?", id)
	if err != nil {
		log.Fatal()
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		log.Fatal()
	}
	return rows
}
