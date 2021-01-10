package db

import (
	"fmt"
	"github.com/xujiajun/nutsdb"
)

/*
	指定key和bucket存入对应数据，bucket类似于表名
*/

func Add(key, val []byte, bucket string) {
	//获取数据库对象
	opt := nutsdb.DefaultOptions

	//指定打开数据库位置，若无数据库则创建
	opt.Dir = "./nutsdb"

	//打开数据库
	db, _ := nutsdb.Open(opt)

	//关闭数据库
	defer db.Close()

	//更新数据库
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			//使用put方法插入数据，若存中相同的key，则更新
			if err := tx.Put(bucket, key, val, 0); err != nil {
				fmt.Println(err)
				return err
			}
			return nil
		}); err != nil {
		fmt.Println(err)
	}
}
