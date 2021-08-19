package models

import (
	"coffee_backend/db"
	"log"
)

type Sequence struct {
	Contrast_no string `json:"ContrastNo" form:"ContrastNo"`
	Before_no   int    `json:"beforeNo" form:"beforeNo"`
	Next_no     int    `json:"nextNo" form:"nextNo"`
	Seq_Range   int    `json:"seqRange" form:"seqRange"`
	Description string `json:"description" form:"description"`
}

//查询一條seq
func (s *Sequence) GetRow() (sequence Sequence, err error) {
	sequence = Sequence{}
	err = db.SqlDB.QueryRow("select before_no,next_no,seq_range from SEQUENCE where contrast_no = ?", s.Contrast_no).
		Scan(&sequence.Before_no, &sequence.Next_no, &sequence.Seq_Range)
	return
}

//將查出來的seq 不管是否有使用都要update
func (sequence *Sequence) UpdateSeq() int64 {
	rs, err := db.SqlDB.Exec("update SEQUENCE set before_no=?,next_no = ? where contrast_no = ?", sequence.Before_no, sequence.Next_no, sequence.Contrast_no)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	// err := db.SqlDB.Where("contrast_no = ? =?", sequence.Contrast_no).Updates(map[string]interface{}{"before_no": sequence.Before_no, "next_no": sequence.Next_no}).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return rows
}
