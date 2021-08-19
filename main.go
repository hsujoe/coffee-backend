package main

import (
	"coffee_backend/db"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":9000")
	// openHttpListen()
}

// func openHttpListen() {
// 	http.HandleFunc("/", receiveClientRequest)
// 	fmt.Println("go server start running...")

// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

// func receiveClientRequest(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
// 	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
// 	w.Header().Set("content-type", "application/json")             //返回数据格式是json

// 	r.ParseForm()
// 	fmt.Println("收到客户端请求: ", r.Form)
// }
