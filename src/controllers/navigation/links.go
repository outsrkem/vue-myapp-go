package navigation

import (
	"github.com/gin-gonic/gin"
	"mana/src/config"
	"mana/src/models"
	"net/http"
	"regexp"
)

var log = config.Log()

// 获取导航链接
func GetResourceLinks(c *gin.Context) {
	id := c.Param("id")

	// id 为空代表请求中没有传递id，则查询所有
	if id == "" {
		res, err := models.FindByResourceLinks()
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
func AddResourceLink(c *gin.Context)  {
	c.JSON(http.StatusOK, "AddResourceLinks")
}

// 修改更新
func UpdateResourceLink(c *gin.Context)  {
	c.JSON(http.StatusOK, "UpdateResourceLink")
}

// 删除
func DeleteResourceLink(c *gin.Context)  {
	c.JSON(http.StatusOK, "DeleteResourceLink")
}