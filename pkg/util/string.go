package util

import (
	"github.com/sony/sonyflake"
	"strconv"
	"time"
)

var st *sonyflake.Sonyflake

func init() {
	t, err := time.Parse("2006-01-02", "2021-10-01")
	if err != nil {
		panic(err)
	}
	st = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime:      t,
		MachineID:      nil,
		CheckMachineID: nil,
	})
}

func GetUniqueIdString() string {
	id, err := st.NextID()
	if err != nil {
		panic(err)
	}
	return strconv.FormatUint(id, 36)
}
