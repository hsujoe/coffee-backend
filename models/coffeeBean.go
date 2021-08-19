package models

import (
	"coffee_backend/db"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/fatih/structs"
)

type CoffeeBean struct {
	Id           int       `json:"id" form:"id"`
	Store_Id     int       `json:"sid" form:"sid"`
	SName        string    `json:"sname" form:"sname"`
	Buy_date     time.Time `json:"buyDate" form:"buyDate"`
	Weight       int       `json:"weight" form:"weight"`
	Price        int       `json:"price" form:"price"`
	Origin       string    `json:"origin" form:"origin"`
	Manor        string    `json:"manor" form:"manor"`
	Name         string    `json:"name" form:"name"`
	Fermentation string    `json:"fermentation" form:"fermentation"`
	Roast        string    `json:"roast" form:"roast"`
	Acidity      string    `json:"acidity" form:"acidity"`
	Flavor       string    `json:"flavor" form:"flavor"`
	Tasting      string    `json:"tasting" form:"tasting"`
	Description  string    `json:"description" form:"description"`
	State        int       `json:"state" form:"state"`
	Image        string    `json:"image" form:"image"`
}

const (
	layoutFromISO = "2006-01-02"
)

//查询所有记录
func (coffeeBean *CoffeeBean) GetCoffeeBeans(name string) (coffeeBeans []CoffeeBean, err error) {
	m := map[string]interface{}{"name": name}
	var where []string
	for _, k := range []string{"name"} {
		if v, ok := m[k]; ok {
			if len(fmt.Sprintf("%v", v)) > 0 {
				where = append(where, fmt.Sprintf("%s like 'vb%svb'", k, fmt.Sprintf("%v", v)))
			}
		}
	}
	sql := "select c.id as id, c.name as name,c.store_id, s.name as Sname,origin,manor,fermentation,roast," +
		"acidity,flavor,tasting,description,c.state as state,c.image, buy_date, weight,price " +
		" from COFFEEBEAN c LEFT JOIN STORE s on s.id = c.store_id where c.state in (0,1) "

	if len(where) > 0 {
		sql += strings.Join(where, " AND ")
		sql = strings.ReplaceAll(sql, "vb", "%")
	}

	fmt.Print(sql)
	// rows, err := db.SqlDB.Query("select id,product_no,name from PRODUCTS where " + strings.Join(where, " AND "))
	rows, err := db.SqlDB.Query(sql)
	if err == nil {
		db.SqlDB.Begin()
		for rows.Next() {
			cb := CoffeeBean{}
			err := rows.Scan(&cb.Id, &cb.Name, &cb.Store_Id, &cb.SName, &cb.Origin, &cb.Manor,
				&cb.Fermentation, &cb.Roast, &cb.Acidity, &cb.Flavor, &cb.Tasting, &cb.Description,
				&cb.State, &cb.Image, &cb.Buy_date, &cb.Weight, &cb.Price)
			if err != nil {
				log.Fatal(err)
			}
			coffeeBeans = append(coffeeBeans, cb)
		}

		rows.Close()
	}
	return
}

//查询一条记录5
func (cb *CoffeeBean) GetCoffeeBean() (coffeeBean CoffeeBean, err error) {
	coffeeBean = CoffeeBean{}
	err = db.SqlDB.QueryRow("select c.id as id, c.name as name,s.id as Store_Id, origin,manor,fermentation,roast,"+
		"acidity,flavor,tasting,description,c.state as state,c.image as image,buy_date, weight,price"+
		" from COFFEEBEAN c LEFT JOIN STORE s on s.id = c.store_id where c.id = ?", cb.Id).
		Scan(&coffeeBean.Id, &coffeeBean.Name, &coffeeBean.Store_Id, &coffeeBean.Origin, &coffeeBean.Manor, &coffeeBean.Fermentation, &coffeeBean.Roast, &coffeeBean.Acidity,
			&coffeeBean.Flavor, &coffeeBean.Tasting, &coffeeBean.Description, &coffeeBean.State, &coffeeBean.Image, &coffeeBean.Buy_date, &coffeeBean.Weight, &coffeeBean.Price)
	return
}

//插入
func (coffeeBean CoffeeBean) CreateCoffeeBean() int64 {
	buyDate := coffeeBean.Buy_date.Format(layoutFromISO)
	t := reflect.TypeOf(coffeeBean)
	va := reflect.ValueOf(coffeeBean)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = va.Field(i).Interface()
	}
	var fieldData []string
	var valesData []string
	for k, _ := range data {
		if k != "SName" {
			fieldData = append(fieldData, k)
		}
	}

	for _, k := range fieldData {
		if v, ok := data[k]; ok {
			if len(fmt.Sprintf("\"%v\"", v)) > 0 {
				if k == "Id" || k == "State" || k == "Weight" || k == "Price" || k == "Store_Id" {
					valesData = append(valesData, fmt.Sprintf("%s", fmt.Sprintf("%v", v)))
				} else if k == "Buy_date" {
					valesData = append(valesData, fmt.Sprintf("\"%v\"", buyDate))
				} else {
					if k != "SName" {
						valesData = append(valesData, fmt.Sprintf("\"%v\"", v))
					}
				}

			}
		}
	}
	// else if k == "Image" {
	// 	s := biu.BytesToBinaryString(coffeeBean.Image)
	// 	valesData = append(valesData, fmt.Sprintf("\"%v\"", s))
	// }
	sql := "insert into COFFEEBEAN ( " + strings.Join(fieldData, ",") + " ) values ( " + strings.Join(valesData, ",") + " )"
	// fmt.Print(sql)
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
func (cb *CoffeeBean) UpdateCoffeeBean() int64 {
	buyDate := cb.Buy_date.Format(layoutFromISO)
	mainMap := structs.Map(cb)
	var mainValesData []string
	for key, value := range mainMap {
		if len(fmt.Sprintf("%v", value)) > 0 {
			if strings.Compare("Id", fmt.Sprintf("%v", key)) != 0 {
				switch key {
				case "State":
					mainValesData = append(mainValesData, fmt.Sprintf("%s = %s ", key, fmt.Sprintf("%v", value)))
				case "Weight":
					mainValesData = append(mainValesData, fmt.Sprintf("%s = %s ", key, fmt.Sprintf("%v", value)))
				case "Price":
					mainValesData = append(mainValesData, fmt.Sprintf("%s = %s ", key, fmt.Sprintf("%v", value)))
				case "Store_Id":
					mainValesData = append(mainValesData, fmt.Sprintf("%s = %s ", key, fmt.Sprintf("%v", value)))
				case "Buy_date":
					mainValesData = append(mainValesData, fmt.Sprintf("%s = \"%s\" ", key, fmt.Sprintf("%v", buyDate)))
				default:
					mainValesData = append(mainValesData, fmt.Sprintf("%s = \"%s\" ", key, fmt.Sprintf("%v", value)))
				}
			}
		}
	}

	tx, _ := db.SqlDB.Begin()
	_, err := db.SqlDB.Exec("update COFFEEBEAN set "+strings.Join(mainValesData, " , ")+" where id = ?", cb.Id)
	if err != nil {
		defer tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
	return 0

}

//删除一条记录
func DeleteCoffeeBean(id int) int64 {
	rs, err := db.SqlDB.Exec("delete from COFFEEBEAN where id = ?", id)
	if err != nil {
		log.Fatal()
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		log.Fatal()
	}
	return rows
}
