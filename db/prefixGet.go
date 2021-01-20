package db

import (
	"fmt"
	"github.com/xujiajun/nutsdb"
)

/*
	前缀查询
*/
func PrefixGet(bucket, key string) nutsdb.Entries {

	opt := nutsdb.DefaultOptions
	opt.Dir = "./nutsdb"
	db, _ := nutsdb.Open(opt)
	defer db.Close()

	var data nutsdb.Entries
	if err := db.View(
		func(tx *nutsdb.Tx) error {
			prefix := []byte(key)
			// 从offset=0开始 ，限制 100 entries 返回
			if entries, err := tx.PrefixScan(bucket, prefix, 5); err != nil {
				return err
			} else {
				data = entries
			}
			return nil
		}); err != nil {
		fmt.Println(err)
	}
	return data
}
