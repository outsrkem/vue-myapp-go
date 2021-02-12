package navigation

import (
	"encoding/json"
	"mana/src/config"
	"mana/src/filters/util"
	"mana/src/models"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

var log = config.Log()

// 获取导航链接
func GetResourceLinks(c *gin.Context) {
	id := c.Param("id")
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	// 条件查询
	activate := c.DefaultQuery("activate", "")
	category := c.DefaultQuery("category", "")

	// id 为空代表请求中没有传递id，则查询所有
	if id == "" {
		res, err := models.FindByResourceLinks(pageSize, page, activate, category)
		if err != nil {
			msg := models.NewResMessage("500", "Query database failed")
			c.JSON(http.StatusInternalServerError, msg)
			return
		}

		msg := models.NewResMessage("200", "successfully")
		returns := models.NewReturns(res, msg)
		c.JSON(http.StatusOK, returns)
		return
	}

	// id为查询的具体记录
	matched, err := regexp.MatchString("^([1-9][0-9]{0,2})$", id)
	if err != nil {
		msg := models.NewResMessage("500", "system exception")
		c.JSON(http.StatusInternalServerError, msg)
		log.Error("GetResourceLinks system exception")
		return
	}

	// 参数校验
	if !matched {
		msg := models.NewResMessage("400", "The parameter ID must be an integer, ^([1-9][0-9]{0,2})$ ")
		c.JSON(http.StatusBadRequest, msg)
		log.Error("GetResourceLinks Query parameter exception, id: ", id)
		return
	}

	// 数据查询
	res, err := models.FindByResourceLinksTheId(id)
	if err != nil {
		msg := models.NewResMessage("500", "Query database failed")
		c.JSON(http.StatusInternalServerError, msg)
		return
	}
	msg := models.NewResMessage("200", "successfully")
	returns := models.NewReturns(res, msg)
	c.JSON(http.StatusOK, returns)
}

// 添加
func AddResourceLink(c *gin.Context) {
	var l = models.NewResourceLinks()

	// 获取token中的userid
	uid := c.MustGet("uid").(string)

	link := make(map[string]interface{})
	c.ShouldBind(&link)
	data, err := json.Marshal(link)
	if err != nil {
		msg := models.NewResMessage("406", "JSON serialization error.")
		c.JSON(http.StatusNotAcceptable, &msg)
		return
	}

	name := gjson.Get(string(data), "name").String()
	url := gjson.Get(string(data), "url").String()
	// 参数校验
	if !util.RegexpMatchString(url, "^https?://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]$") {
		msg := models.NewResMessage("400", "The 'url' parameter is abnormal, ^https?://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]$ ")
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	activate := gjson.Get(string(data), "activate").String()
	if !util.RegexpMatchString(activate, "^(0|1)$") {
		msg := models.NewResMessage("400", "The 'activate' is abnormal, ^(0|1)$ ")
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	category := gjson.Get(string(data), "category").String()
	if !util.RegexpMatchString(category, "^[1-5]$") {
		msg := models.NewResMessage("400", "The 'category' is abnormal, ^[1-5]$ ")
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	describes := gjson.Get(string(data), "describes").String()
	if uid == "" || name == "" || url == "" || activate == "" || category == "" || describes == "" {
		msg := models.NewResMessage("400", "The request body is abnormal. The key field cannot be empty.")
		c.JSON(http.StatusBadRequest, &msg)
		return
	}

	l.USERID = uid
	l.LINKNAME = name
	l.LINKURL = url
	l.ACTIVATE = activate
	l.CATEGORY = category
	l.DESCRIBES = describes

	// 插入到数据库
	if _, err = models.InsertResourceLink(l); err != nil {
		msg := models.NewResMessage("406", "Insert error .")
		c.JSON(http.StatusNotAcceptable, &msg)
		return
	}
	msg := models.NewResMessage("201", "Successful.")
	c.JSON(http.StatusCreated, &msg)
}

// 修改更新
func UpdateResourceLink(c *gin.Context) {
	var l = models.NewResourceLinks()

	id := c.Param("id")
	// id为查询的具体记录，用于删除
	matched, err := regexp.MatchString("^([1-9][0-9]{0,2})$", id)
	if err != nil {
		msg := models.NewResMessage("500", "system exception")
		c.JSON(http.StatusInternalServerError, msg)
		log.Error("GetResourceLinks system exception")
		return
	}

	// 参数校验
	if !matched {
		msg := models.NewResMessage("400", "The parameter ID must be an integer, ^([1-9][0-9]{0,2})$ ")
		c.JSON(http.StatusBadRequest, msg)
		log.Error("GetResourceLinks Query parameter exception, id: ", id)
		return
	}

	link := make(map[string]interface{})
	c.ShouldBind(&link)
	data, err := json.Marshal(link)
	if err != nil {
		msg := models.NewResMessage("406", "JSON serialization error.")
		c.JSON(http.StatusNotAcceptable, &msg)
		return
	}

	name := gjson.Get(string(data), "name").String()
	url := gjson.Get(string(data), "url").String()
	activate := gjson.Get(string(data), "activate").String()
	category := gjson.Get(string(data), "category").String()
	describes := gjson.Get(string(data), "describes").String()

	if name == "" || url == "" || category == "" || describes == "" || activate == "" || id == "" {
		msg := models.NewResMessage("400", "The request body is abnormal. The key field cannot be empty.")
		c.JSON(http.StatusBadRequest, &msg)
		return
	}

	l.ID = id
	l.LINKNAME = name
	l.LINKURL = url
	l.ACTIVATE = activate
	l.CATEGORY = category
	l.DESCRIBES = describes

	// 插入到数据库
	n, err := models.UpdateResourceLinkToDb(l)
	if err != nil {
		msg := models.NewResMessage("406", "Update error .")
		c.JSON(http.StatusNotAcceptable, &msg)
		return
	}
	if n == "0" {
		msg := models.NewResMessage("304", "Affected row: 0.")
		c.JSON(http.StatusNotModified, &msg)
		return
	}

	msg := models.NewResMessage("200", "Successful.")
	c.JSON(http.StatusOK, &msg)
}

// 删除
func DeleteResourceLink(c *gin.Context) {
	id := c.Param("id")
	// id为查询的具体记录，用于删除
	matched, err := regexp.MatchString("^([1-9][0-9]{0,2})$", id)
	if err != nil {
		msg := models.NewResMessage("500", "system exception")
		c.JSON(http.StatusInternalServerError, msg)
		log.Error("GetResourceLinks system exception")
		return
	}

	// 参数校验
	if !matched {
		msg := models.NewResMessage("400", "The parameter ID must be an integer, ^([1-9][0-9]{0,2})$ ")
		c.JSON(http.StatusBadRequest, msg)
		log.Error("GetResourceLinks Query parameter exception, id: ", id)
		return
	}

	// 删除记录
	n, err := models.DeleteLink(id)
	if err != nil {
		msg := models.NewResMessage("406", "Update error .")
		c.JSON(http.StatusNotAcceptable, &msg)
		return
	}
	if n == "0" {
		msg := models.NewResMessage("304", "Affected row: 0.")
		c.JSON(http.StatusNotModified, &msg)
		return
	}

	msg := models.NewResMessage("200", "Successful.")
	c.JSON(http.StatusOK, &msg)
}
