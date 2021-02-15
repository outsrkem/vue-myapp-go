package resource

import (
	"fmt"
	"mana/src/filters/util"
	"mana/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// GetHostMonitorInfo 获取监控信息
func GetHostMonitorInfo(c *gin.Context) {
	url := `http://10.10.10.10:9090/api/v1/query?query=100+-+%28%28node_memory_MemFree_bytes%2Bnode_memory_Cached_bytes%2Bnode_memory_Buffers_bytes%29+%2F+node_memory_MemTotal_bytes+*+100%29&time=1613308876`
	data := util.GetMemoryAvg(url)

	array := gjson.Get(string(*data), "data.result").Array()
	fmt.Println(array)
	num := len(array)
	test := make(map[interface{}]string, num)
	for i := 0; i < num; i++ {
		test[gjson.Get(string(*data), "result").Array()] = array[i].Str
	}
	fmt.Println(test)
	// 临时返回
	msg := models.NewResMessage("200", "successfully")
	pageInfo := models.NewPageInfo(1, 1, 1, 1)
	items := test
	res := models.NewResponse(items, pageInfo)
	returns := models.NewReturns(res, msg)
	c.JSON(http.StatusOK, &returns)
}
