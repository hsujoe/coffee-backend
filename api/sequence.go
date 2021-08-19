package api

import (
	. "coffee_backend/models"
	"fmt"
)

//获得一seq 並新重更新
func GetSeq(seqType string) (newSeq int) {
	s := Sequence{
		Contrast_no: seqType,
	}
	rs, _ := s.GetRow()

	updates := Sequence{
		Before_no:   rs.Next_no,
		Next_no:     rs.Next_no + rs.Seq_Range,
		Contrast_no: seqType,
	}
	row := updates.UpdateSeq()
	fmt.Sprintf("update successful %d", row)
	return rs.Next_no
}
