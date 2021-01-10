package db

import (
	"fmt"
	"github.com/xujiajun/nutsdb"
)

/*
	指定key和bucket获取数据
*/
func Get(key []byte, bucket string) *nutsdb.Entry {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nutsdb"
	db, _ := nutsdb.Open(opt)
	defer db.Close()

	//var list []byte
	var list *nutsdb.Entry

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			if e, err := tx.Get(bucket, key); err != nil {
				return nil
			} else {
				list = e
				//fmt.Println(string(e.Value))
			}
			return nil
		}); err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}
	return list
}
