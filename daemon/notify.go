package daemon

import (
	"birthday/models/bmob"
	"birthday/request"
	"encoding/json"
	"fmt"
	"time"
)

func Notify() {

	for i := 0; i < 4; i++ {
		users, err := bmob.UserList()
		if err != nil {
			fmt.Println("user list err: ", err)
			continue
		}
		list, err := bmob.BirthdayList()
		if err != nil {
			fmt.Println("birtday list err: ", err)
			continue
		}

		mUser := bmob.UserListToMap(users)
		for _, item := range list {
			birthday, err := time.ParseInLocation("2006-01-02 15:04:05", item.SolarCalendar.Date, time.Now().Location())
			if err != nil {
				fmt.Println("parse in location: ", err)
				continue
			}
			//fmt.Println("birthday: ", birthday.Day(), birthday.Month(), )
			if birthday.Month() == time.Now().Month() && birthday.Day() == time.Now().Day() {
				var own *bmob.User
				if v, ok := mUser[int(item.Own)]; ok {
					own = v
				} else {
					continue
				}
				//phone := strconv.FormatFloat(own.Phone, 'f', -1, 64)
				if item.SendSmsDate != time.Now().Add(48*time.Hour).Format("20060102") { //判断是否发过短信
					r, err := sendSms(own.Phone, fmt.Sprintf("亲爱的%s，后天是%s的生日，记得买礼物或者发送祝福喔，本短信来自小程序生日工具。", own.UserName, item.Name))
					if err != nil {
						fmt.Println("send sms err:", err)
						continue
					}
					//{"smsId":39526947}
					var smsRes map[string]int
					if err := json.Unmarshal(r, &smsRes); err != nil {
						continue
					}
					if v, ok := smsRes["smsId"]; ok && v > 0 {
						err = item.UpdateSendSmsDate(time.Now().Format("20060102"))
						if err != nil {
							fmt.Println("updatesend sms date err: ", err)
						}
					}

				}
			}
		}
		time.Sleep(time.Second * 60)
	}

}

func sendSms(phone string, content string) (r []byte, err error) {
	data := map[string]string{
		"mobilePhoneNumber": phone,
		"template":          "srtx",
		"content":           content,
	}
	r, err = request.Post("https://api.bmob.cn/1/requestSms", data)
	if err != nil {
		return
	}
	return
}
