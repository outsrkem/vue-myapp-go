package db

import (
	"fmt"
	"github.com/xujiajun/nutsdb"
	"log"
)

/*
	指定key和bucket获取数据
*/
func Get(key []byte, bucket string) *nutsdb.Entry {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nutsdb"
	db, _ := nutsdb.Open(opt)
	defer db.Close()

	var list *nutsdb.Entry

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			e, err := tx.Get(bucket, key)
			if err != nil {
				fmt.Println(err)
				return err
			}
			list = e
			return nil
		}); err != nil {
		log.Println(err)
	}
	return list
}
