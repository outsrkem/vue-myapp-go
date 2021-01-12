package impl

import (
	"fmt"
	"github.com/xujiajun/nutsdb"
)

func BackFile(onOff bool) {
	if onOff {
		opt := nutsdb.DefaultOptions
		opt.Dir = "./nutsdb"
		db, _ := nutsdb.Open(opt)

		err := db.Backup("./backup")
		if err != nil {
			fmt.Println("备份错误", err)
		}
		db.Close()
	}
}
