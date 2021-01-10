package db

import (
	"fmt"
	"github.com/xujiajun/nutsdb"
)

/*
	获取指定bucket的数据总数
*/
func GetIndex(bucket string) int {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nutsdb"
	db, _ := nutsdb.Open(opt)
	defer db.Close()

	index := 0

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			if e, err := tx.GetAll(bucket); err != nil {
				return nil
			} else {
				index = len(e)
			}
			return nil
		}); err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}

	return index
}
