package impl

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"menu/db"
	"strings"
	"time"
)

/*
	获取所有服务器列表性能数据
*/

func (l *LinuxList) GetAll() {
	//查询服务器，获取数据索引和账号信息
	index := db.GetIndex(db.LinuxList)
	data := db.GetAll(db.LinuxList)
	if data == nil {
		l.MetaInfo.Status = 502
		l.MetaInfo.Msg = "GetAll获取数据库错误"
		l.MetaInfo.RequestTime = time.Now().Unix()
		return
	}

	l.MetaInfo.Status = 200
	l.MetaInfo.Msg = "success"
	l.MetaInfo.RequestTime = time.Now().Unix()

	l.Response.PageInfo.Page = 1
	l.Response.PageInfo.PageNum = 1
	l.Response.PageInfo.PageSize = 1
	l.Response.PageInfo.Total = 1

	//创建ItemsL对象，使用数据库索引作为长度
	item := make([]ItemsL, index)
	//创建数据库对应的结构体对象
	var linuxuser LinuxListUser
	for i, k := range data {
		//反序列化数据库数据
		err := json.Unmarshal(k.Value, &linuxuser)
		if err != nil {
			fmt.Println(err)
			return
		}
		//根据数据库账号信息，获取对应服务器性能
		linuxcmd, err := GetLinuxCmd(&linuxuser)
		if err != nil {
			fmt.Println(err)
			return
		}

		//数据库账号信息数字符串，需要处理获取对应数据
		linux := strings.Fields(linuxcmd)
		item[i].Hostname = linux[3]
		item[i].HostIp = linuxuser.Ip
		item[i].Id = i + 1

		//处理服务器资源切片，赋值给新的切片,a切片是资源使用情况，b切片是使用比例
		a := []int{}
		b := []int{}
		for x := 0; x < 3; x++ {
			s := strings.Split(strings.Fields(linuxcmd)[x], ":")
			a = append(a, cast.ToInt(s[0]), cast.ToInt(s[1]))
			b = append(b, cast.ToInt(cast.ToFloat64(s[1])/cast.ToFloat64(s[0])*100))
		}

		item[i].Memory.Total = a[0]
		item[i].Memory.Used = a[1]
		item[i].Memory.Section = 666
		item[i].Memory.Proportion = b[0]

		item[i].Cpu.Total = a[2]
		item[i].Cpu.Used = a[3]
		item[i].Cpu.Section = 666
		item[i].Cpu.Proportion = b[1]

		item[i].Disk.Total = a[4]
		item[i].Disk.Used = a[5]
		item[i].Disk.Section = 666
		item[i].Disk.Proportion = b[2]
	}
	l.Response.Items = item
}
