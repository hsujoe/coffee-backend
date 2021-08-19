package models

import (
	"coffee_backend/db"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

type Brew struct {
	Id               int     `json:"id" form:"id"`
	Bean_Id          int     `json:"bid" form:"bid"`
	BName            string  `json:"bname" form:"bname"`
	Grinder          string  `json:"grinder" form:"grinder"`
	Scale            float32 `json:"scale" form:"scale"`
	Steaming         int     `json:"steaming" form:"steaming"`
	WaterCuts        int     `json:"waterCuts" form:"waterCuts"`
	Totalwater       int     `json:"totalWater" form:"totalWater"`
	Watertemperature int     `json:"waterTemperature" form:"waterTemperature"`
	Description      string  `json:"description" form:"description"`
}

//查询所有记录
func (brew *Brew) GetBrews(name string) (Brews []Brew, err error) {
	m := map[string]interface{}{"name": name}
	var where []string
	for _, k := range []string{"name"} {
		if v, ok := m[k]; ok {
			if len(fmt.Sprintf("%v", v)) > 0 {
				where = append(where, fmt.Sprintf("%s like 'vb%svb'", k, fmt.Sprintf("%v", v)))
			}
		}
	}
	sql := "select b.id as id, b.bean_id, c.name as bname,grinder,scale,steaming,watercuts,watertemperature,totalwater,b.description " +
		"from BREW b LEFT JOIN COFFEEBEAN c on c.id = b.bean_id"

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
			brew := Brew{}
			err := rows.Scan(&brew.Id, &brew.Bean_Id, &brew.BName, &brew.Grinder,
				&brew.Scale, &brew.Steaming, &brew.WaterCuts, &brew.Watertemperature,
				&brew.Totalwater, &brew.Description)
			if err != nil {
				log.Fatal(err)
			}
			Brews = append(Brews, brew)
		}

		rows.Close()
	}
	return
}

//查询一条记录5
func (b *Brew) GetBrew() (brew Brew, err error) {
	brew = Brew{}
	err = db.SqlDB.QueryRow("select b.id as id, b.bean_id, c.name as bname,grinder,scale,steaming,watercuts,watertemperature,totalwater,b.description "+
		"from BREW b LEFT JOIN COFFEEBEAN c on c.id = b.bean_id where b.id = ?", b.Id).
		Scan(&brew.Id, &brew.Bean_Id, &brew.BName, &brew.Grinder,
			&brew.Scale, &brew.Steaming, &brew.WaterCuts, &brew.Watertemperature, &brew.Totalwater, &brew.Description)
	return
}

//插入
func (brew Brew) CreateBrew() int64 {
	t := reflect.TypeOf(brew)
	va := reflect.ValueOf(brew)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = va.Field(i).Interface()
	}
	var fieldData []string
	var valesData []string
	for k, _ := range data {
		if k != "BName" {
			fieldData = append(fieldData, k)
		}
	}

	for _, k := range fieldData {
		if v, ok := data[k]; ok {
			if len(fmt.Sprintf("\"%v\"", v)) > 0 {
				if k == "Id" || k == "State" || k == "Scale" ||
					k == "Steaming" || k == "WaterCuts" || k == "Totalwater" ||
					k == "Watertemperature" {
					valesData = append(valesData, fmt.Sprintf("%s", fmt.Sprintf("%v", v)))
				} else {
					if k != "BName" {
						valesData = append(valesData, fmt.Sprintf("\"%v\"", v))
					}
				}

			}
		}
	}

	sql := "insert into BREW ( " + strings.Join(fieldData, ",") + " ) values ( " + strings.Join(valesData, ",") + " )"
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
func (brew *Brew) UpdateBrew() int64 {
	mainMap := structs.Map(brew)
	var mainValesData []string
	for key, value := range mainMap {
		if len(fmt.Sprintf("%v", value)) > 0 {
			if strings.Compare("Id", fmt.Sprintf("%v", key)) != 0 {
				mainValesData = append(mainValesData, fmt.Sprintf("%s = \"%v\" ", key, fmt.Sprintf("%v", value)))
			}
		}
	}

	tx, _ := db.SqlDB.Begin()
	_, err := db.SqlDB.Exec("update BREW set "+strings.Join(mainValesData, " , ")+" where id = ?", brew.Id)
	if err != nil {
		defer tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
	return 0

}

//删除一条记录
func DeleteBrew(id int) int64 {
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
