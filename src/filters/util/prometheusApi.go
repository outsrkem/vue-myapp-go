package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GetMemoryAvg 获取平均内存使用率
func GetMemoryAvg(url string) *[]byte {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Get error:", err)
		return nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return &body
}
