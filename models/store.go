package models

import (
	"coffee_backend/db"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

type Store struct {
	Id      int    `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Tel     string `json:"tel" form:"tel"`
	Address string `json:"address" form:"address"`
	Page    string `json:"page" form:"page"`
	Fb      string `json:"fb" form:"fb"`
	Ig      string `json:"ig" form:"ig"`
	Memo    string `json:"memo" form:"memo"`
	State   int    `json:"state" form:"state"`
	Image   string `json:"image" form:"image"`
}

//查询所有记录
func (store *Store) GetStores(name string) (stores []Store, err error) {
	m := map[string]interface{}{"name": name}
	var where []string
	for _, k := range []string{"name"} {
		if v, ok := m[k]; ok {
			if len(fmt.Sprintf("%v", v)) > 0 {
				where = append(where, fmt.Sprintf("%s like 'vb%svb'", k, fmt.Sprintf("%v", v)))
			}
		}
	}
	sql := "select id,name,tel,address,page,fb,ig,memo,state,image  from STORE where state in (0,1) "
	if len(where) > 0 {
		sql += strings.Join(where, " AND ")
		sql = strings.ReplaceAll(sql, "vb", "%")
	}

	fmt.Print(sql)
	// rows, err := db.SqlDB.Query("select id,product_no,name from PRODUCTS where " + strings.Join(where, " AND "))
	rows, err := db.SqlDB.Query(sql)
	db.SqlDB.Begin()
	for rows.Next() {
		store := Store{}
		// err := rows.Scan(&store.Id, &store.Name, &store.Tel, &store.Address, &store.State, &store.Fb, &store.Ig, &store.Memo, &store.State)
		err := rows.Scan(&store.Id, &store.Name, &store.Tel, &store.Address, &store.Page, &store.Fb, &store.Ig, &store.Memo, &store.State, &store.Image)
		if err != nil {
			log.Fatal(err)
		}
		stores = append(stores, store)
	}
	rows.Close()
	return
}

//查询推薦
func (store *Store) GetRecommendStores() (stores []Store, err error) {
	sql := "select id,name,image  from STORE where state = 2"

	fmt.Print(sql)
	// rows, err := db.SqlDB.Query("select id,product_no,name from PRODUCTS where " + strings.Join(where, " AND "))
	rows, err := db.SqlDB.Query(sql)
	db.SqlDB.Begin()
	for rows.Next() {
		store := Store{}
		err := rows.Scan(&store.Id, &store.Name, &store.Image)
		if err != nil {
			log.Fatal(err)
		}
		stores = append(stores, store)
	}
	rows.Close()
	return
}

//查询一条记录5
func (s *Store) GetStore() (store Store, err error) {
	store = Store{}
	err = db.SqlDB.QueryRow("select id,name,tel,address,page,fb,ig,memo,state,image from STORE where id = ?", s.Id).
		Scan(&store.Id, &store.Name, &store.Tel, &store.Address, &store.Page, &store.Fb, &store.Ig, &store.Memo, &store.State, &store.Image)
	return
}

//插入
func (store Store) CreateStore() int64 {
	t := reflect.TypeOf(store)
	va := reflect.ValueOf(store)
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
				if k == "Id" || k == "State" {
					valesData = append(valesData, fmt.Sprintf("%s", fmt.Sprintf("%v", v)))
				} else {
					valesData = append(valesData, fmt.Sprintf("\"%v\"", v))
				}

			}
		}
	}

	sql := "insert into STORE ( " + strings.Join(fieldData, ",") + " ) values ( " + strings.Join(valesData, ",") + " )"
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
func (store *Store) UpdateStore() int64 {
	mainMap := structs.Map(store)
	var mainValesData []string
	for key, value := range mainMap {
		if len(fmt.Sprintf("%v", value)) > 0 {
			if strings.Compare("Id", fmt.Sprintf("%v", key)) != 0 {
				mainValesData = append(mainValesData, fmt.Sprintf("%s = \"%v\" ", key, fmt.Sprintf("%v", value)))
			}
		}
	}

	tx, _ := db.SqlDB.Begin()
	_, err := db.SqlDB.Exec("update STORE set "+strings.Join(mainValesData, " , ")+" where id = ?", store.Id)
	if err != nil {
		defer tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
	return 0

}

//删除一条记录
func DeleteStore(id int) int64 {
	rs, err := db.SqlDB.Exec("delete from STORE where id = ?", id)
	if err != nil {
		log.Fatal()
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		log.Fatal()
	}
	return rows
}
