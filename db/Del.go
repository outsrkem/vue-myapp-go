package db

import (
	"fmt"
	"github.com/xujiajun/nutsdb"
)

func Del(key []byte, bucket string) {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nutsdb"
	db, _ := nutsdb.Open(opt)
	defer db.Close()

	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Delete(bucket, key); err != nil {
				fmt.Println(err)
				return err
			}
			return nil
		}); err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}
}
