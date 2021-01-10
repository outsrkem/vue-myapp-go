package db

import (
	"fmt"
	"github.com/xujiajun/nutsdb"
)

/*
	获取指定的bucket所有数据
*/
func GetAll(bucket string) nutsdb.Entries {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nutsdb"
	db, _ := nutsdb.Open(opt)
	defer db.Close()

	var list nutsdb.Entries

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			if e, err := tx.GetAll(bucket); err != nil {
				return nil
			} else {
				list = e
			}
			return nil
		}); err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}

	return list
}
