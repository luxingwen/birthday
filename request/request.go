package request

import (
	"birthday/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
)

var (
	BmobAppKey  string
	BmobRestKey string
)

func init() {
	BmobAppKey = models.AppConfig.DefaultString("bmob::app_key", "")
	BmobRestKey = models.AppConfig.DefaultString("bmob::rest_key", "")
}
func Get(address string, m map[string]string) ([]byte, error) {
	fmt.Println(BmobAppKey)
	req := httplib.Get(address)
	for k, v := range m {
		req.Param(k, v)
	}
	req.Header("X-Bmob-Application-Id", BmobAppKey)
	req.Header("X-Bmob-REST-API-Key", BmobRestKey)
	req.Header("Content-Type", "application/json")
	return req.Bytes()
}

func Post(address string, m interface{}) ([]byte, error) {
	var data []byte
	if v, ok := m.([]byte); ok {
		data = v
	} else {
		b, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		data = b
	}
	req := httplib.Post(address)
	req.Header("X-Bmob-Application-Id", BmobAppKey)
	req.Header("X-Bmob-REST-API-Key", BmobRestKey)
	req.Header("Content-Type", "application/json")
	req.Body(data)
	return req.Bytes()
}

func Put(address string, m interface{}) ([]byte, error) {
	var data []byte
	if v, ok := m.([]byte); ok {
		data = v
	} else {
		b, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		data = b
	}
	req := httplib.Put(address)
	req.Header("X-Bmob-Application-Id", BmobAppKey)
	req.Header("X-Bmob-REST-API-Key", BmobRestKey)
	req.Header("Content-Type", "application/json")
	req.Body(data)
	return req.Bytes()
}
