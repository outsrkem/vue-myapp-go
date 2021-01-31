package models

import (
	"mana/src/connections/database/mysql"
	"strconv"
	"time"
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

// 查询导航链接列表
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
	// 3 一定要关闭连接
	defer rows.Close()
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

// 根据id查询导航链接
func FindByResourceLinksTheId(id int) (*map[string]string, error) {
	var l resourceLinks
	sqlStr := `SELECT ID,USERID,LINKNAME,LINKURL,DESCRIBES,CATEGORY,ACTIVATE,CREATETIME,UPDATETIME FROM resource_links WHERE ID=?;`
	var row = mysql.DB.QueryRow(sqlStr, id)
	err := row.Scan(&l.ID, &l.USERID, &l.LINKNAME, &l.LINKURL, &l.DESCRIBES, &l.CATEGORY, &l.ACTIVATE,
		&l.CREATETIME, &l.UPDATETIME)

	if err != nil {
		log.Error("exec  query failed, ", sqlStr , "ID=",id , err)
		return nil, err
	}

	createTime, _ := strconv.ParseInt(l.CREATETIME, 10, 64)
	updateTime, _ := strconv.ParseInt(l.UPDATETIME, 10, 64)
	item := make(map[string]string)
	item["id"] = l.ID
	item["name"] = l.LINKNAME
	item["content"] = l.LINKURL
	item["describes"] = l.DESCRIBES
	item["activate"] = l.ACTIVATE
	item["category"] = l.CATEGORY
	item["createTime"] = time.Unix(createTime, 0).Format("2006-01-02 15:04:05")
	item["updateTime"] = time.Unix(updateTime, 0).Format("2006-01-02 15:04:05")

	returns := item
	return &returns, err
}
