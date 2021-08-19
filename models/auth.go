package models

import (
	"coffee_backend/db"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	uuid "github.com/satori/go.uuid"
)

type LoginToken struct {
	jwt.Payload
	// ID       uint   `json:"id"`
	ID       uint   `json:"id"`
	UserId   string `json:"userId"`
	Username string `json:"userName"`
	Password string `json:"password"`
	Image    string `json:"image" form:"image"`
}

var privateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
var publicKey = &privateKey.PublicKey
var hs = jwt.NewES256(
	jwt.ECDSAPublicKey(publicKey),
	jwt.ECDSAPrivateKey(privateKey),
)

//確認帳密
func CheckUser(userId string, passWord string) (loginToken LoginToken, err error) {
	m := map[string]interface{}{"user_Id": userId, "password": passWord}
	var where []string
	for _, k := range []string{"user_Id", "password"} {
		if v, ok := m[k]; ok {
			if len(fmt.Sprintf("%v", v)) > 0 {
				where = append(where, fmt.Sprintf(" %s = '%s' ", fmt.Sprintf("%v", k), fmt.Sprintf("%v", v)))
			}
		}
	}
	sql := "select id, user_id,username,password,image from USERS"
	sql += " where " + strings.Join(where, " AND ")
	// sql := "select id, user_id,username,password from USERS where user_id = " + fmt.Sprintf("%v", userId) + " and password= " + fmt.Sprintf("%v", passWord)
	fmt.Print(sql)
	rows, err := db.SqlDB.Query(sql)
	// var result = false
	if err == nil {
		db.SqlDB.Begin()
		for rows.Next() {
			LT := LoginToken{}
			err := rows.Scan(&LT.ID, &LT.UserId, &LT.Username, &LT.Password, &LT.Image)
			if err != nil {
				log.Fatal(err)
			}
			loginToken = LT
		}

		rows.Close()
	}
	return
}

// 签名
func Sign(userId string, password string) (string, error) {
	now := time.Now()
	pl := LoginToken{
		Payload: jwt.Payload{
			Issuer:         "coolcat",
			Subject:        "login",
			Audience:       jwt.Audience{},
			ExpirationTime: jwt.NumericDate(now.Add(7 * 24 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          uuid.NewV4().String(),
		},
		UserId:   userId,
		Password: password,
	}
	token, err := jwt.Sign(pl, hs)
	return string(token), err
}

// 验证
func Verify(token []byte) (*LoginToken, error) {
	pl := &LoginToken{}
	_, err := jwt.Verify(token, hs, pl)
	return pl, err
}

//查询一条记录5
func (l *LoginToken) GetUser() (loginToken LoginToken, err error) {
	loginToken = LoginToken{}
	err = db.SqlDB.QueryRow("select id, user_id,username,password,image from USERS where id = ?", l.ID).
		Scan(&loginToken.ID, &loginToken.UserId, &loginToken.Username, &loginToken.Password, &loginToken.Image)
	return
}
