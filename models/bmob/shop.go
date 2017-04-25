package bmob

import (
	"birthday/request"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	//"net/url"
	"strconv"
	"strings"
	"time"
)

type Shop struct {
	ObjectId    string   `json:"objectId"`
	Location    string   `json:"location"`
	LocationGeo Location `json:"locationGeo"`
	Trans       bool     `json:"trans"`
}

type Location struct {
	Type      string  `json:"__type"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func ShopList() ([]*Shop, error) {
	count, err := count(getAddress("shopcake2"))
	if err != nil {
		return nil, err
	}
	var r []*Shop
	for i := 0; i < count; i += 100 {

		v, err := ShopOnPage("100", strconv.Itoa(i))
		if err != nil {
			fmt.Println("birthday on page error: ", err)
			continue
		}
		r = append(r, v...)
	}
	return r, nil
}

func ShopOnPage(limit, skip string) ([]*Shop, error) {
	m := map[string]string{"limit": limit, "skip": skip}
	b, err := request.Get(getAddress("shopcake2"), m)
	if err != nil {
		return nil, err
	}

	//fmt.Println("body:" + string(b))
	var res struct {
		Data []*Shop `json:"results"`
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println("err--->", err)
		return nil, err
	}
	var shops []*Shop
	for _, item := range res.Data {
		if item.Trans {
			continue
		}
		shops = append(shops, item)
	}
	return shops, nil
}

type QQLocation struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func Transh() ([]*Shop, error) {
	list, err := ShopList()
	if err != nil {
		return nil, err
	}

	var qqLocations []*QQLocation

	str := ""
	for i, item := range list {
		locs := strings.Split(item.Location, ",")
		s := locs[1] + "," + locs[0] + ";"
		str += s
		if i%30 == 0 {
			str = str[:len(str)-1]
			urls := "http://apis.map.qq.com/ws/coord/v1/translate?key=PP2BZ-RZHCI-MYDGJ-5DDQN-I2BZO-2PFYH&type=3&locations=" + str

			req := httplib.Get(urls)
			body, err := req.Bytes()
			if err != nil {
				fmt.Println("qqq-->err:", err)
				return nil, err
			}
			var res struct {
				Status    int64         `json:"status"`
				Message   string        `json:"message"`
				Locations []*QQLocation `json:"locations"`
			}
			if err := json.Unmarshal(body, &res); err != nil {
				fmt.Println("err--->", err)
				return nil, err
			}

			fmt.Println(string(body))
			qqLocations = append(qqLocations, res.Locations...)
			str = ""
			time.Sleep(time.Second)
		}
	}

	str = str[:len(str)-1]
	urls := "http://apis.map.qq.com/ws/coord/v1/translate?key=PP2BZ-RZHCI-MYDGJ-5DDQN-I2BZO-2PFYH&type=3&locations=" + str

	req := httplib.Get(urls)
	body, err := req.Bytes()
	if err != nil {
		fmt.Println("qqq-->err:", err)
		return nil, err
	}
	fmt.Println(string(body))
	var res struct {
		Status    int64         `json:"status"`
		Message   string        `json:"message"`
		Locations []*QQLocation `json:"locations"`
	}
	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Println("err--->", err)
		return nil, err
	}

	fmt.Println(string(body))
	qqLocations = append(qqLocations, res.Locations...)

	fmt.Println(len(qqLocations), len(list))
	if len(qqLocations) != len(list) {
		fmt.Println("len err")
		return nil, nil
	}
	for i, item := range qqLocations {
		list[i].LocationGeo.Latitude = item.Lat
		list[i].LocationGeo.Longitude = item.Lng
	}
	return list, nil
}

func (this *Shop) Update() error {
	m := map[string]interface{}{"trans": true, "locationGeo": this.LocationGeo}
	r, err := request.Put(getAddress("shopcake2")+"/"+this.ObjectId, m)
	if err != nil {
		return err
	}
	fmt.Println(string(r))
	return nil
}
