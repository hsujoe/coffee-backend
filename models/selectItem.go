package models

import (
	"coffee_backend/db"
	"fmt"
	"log"
)

type SelectItem struct {
	Id   int    `json:"value" form:"id"`
	Name string `json:"text" form:"name"`
}

//查询所有记录
func GetSelectItem(dataName string) (selectItems []SelectItem, err error) {
	sql := "select id,name from " + dataName + " where state in (0,2)"
	// fmt.Sprintf(sql, "%s", dataName)
	fmt.Print(sql)
	rows, err := db.SqlDB.Query(sql)
	if err == nil {
		db.SqlDB.Begin()
		for rows.Next() {
			item := SelectItem{}
			err := rows.Scan(&item.Id, &item.Name)
			if err != nil {
				log.Fatal(err)
			}
			selectItems = append(selectItems, item)
		}
		rows.Close()
	}
	return selectItems, err
}
