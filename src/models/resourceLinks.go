package models

import (
	"mana/src/connections/database/mysql"
)

type resourceLinks struct {
	ID         string // id
	USERID     string // 用户id
	LINKNAME   string // 链接名
	LINKURL    string // url
	DESCRIBES  string // 说明
	CATEGORY   string // 类别
	ACTIVATE   string // 状态
	CREATETIME string // 创建时间
	UPDATETIME string // 更新时间
}

// 查询导航链接
func FindByResourceLinks() (*map[string]interface{}, error) {
	var l resourceLinks
	var items []map[string]string
	items = make([]map[string]string, 0)
	sqlStr := `SELECT ID,USERID,LINKNAME,LINKURL,DESCRIBES,CATEGORY,ACTIVATE,CREATETIME,UPDATETIME FROM resource_links;`
	rows, err := mysql.DB.Query(sqlStr)
	if err != nil {
		log.Error("exec  query failed, err", sqlStr, err)
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&l.ID, &l.USERID, &l.LINKNAME, &l.LINKURL, &l.DESCRIBES,
			&l.CATEGORY, &l.ACTIVATE, &l.CREATETIME, &l.UPDATETIME)
		if err != nil {
			log.Error("Exception reading database result", err)
		}

		item := make(map[string]string)
		item["id"] = l.ID
		item["name"] = l.LINKNAME
		item["content"] = l.LINKURL
		item["describes"] = l.DESCRIBES
		item["category"] = l.CATEGORY
		items = append(items, item)
	}

	returns := NewResponse(items, nil)
	return &returns, err
}
