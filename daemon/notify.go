package daemon

import (
	"birthday/models"
	"birthday/models/bmob"
	"birthday/request"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	January = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func SendWchatNotify() {

	// formId := "jjjj"
	// err := weichatNotify("ocBzs0I9fkcJ6kq1_xSxtN5dj7bI", "您已添加XXX，于2016-15-45生日，生日工具会在生日的2天前提醒您", time.Now().Add(48*time.Hour).Format("2006-01-02"), formId)
	// if err != nil {
	// 	fmt.Println("weichatnotify err: ", err)
	// }

	users, err := bmob.UserList()
	if err != nil {
		fmt.Println("user list err: ", err)
		return
	}

	mUser := bmob.UserListToMap(users)

	w := map[string]string{"createdAt": "skip"}
	list, err := bmob.BirthdayList(w)
	if err != nil {
		fmt.Println("birtday list err: ", err)

	}

	for _, item := range list {
		birthday, err := time.ParseInLocation("2006-01-02 15:04:05", item.CreatedAt, time.Now().Location())
		if err != nil {
			fmt.Println("parse in location: ", err)
			return
		}

		solarCalendar, err := time.ParseInLocation("2006-01-02 15:04:05", item.SolarCalendar.Date, time.Now().Location())
		if err != nil {
			fmt.Println("parse in location: ", err)
			return
		}

		//取用户信息
		var own *bmob.User
		if v, ok := mUser[int(item.Own)]; ok {
			own = v
		}

		formId := item.FormId
		msg := fmt.Sprintf("您已添加%s生日，生日工具会在%s生日的2天前提醒您", item.Name, item.Name)
		err = weichatNotify(own.Openid, msg, time.Now().Add(48*time.Hour).Format("2006-01-02"), formId)
		if err != nil {
			fmt.Println("weichatnotify err: ", err)
			return
		}

		fmt.Println("birthday", birthday, solarCalendar, msg)
		fmt.Println("item.CreatedAt", item.Name, item.SolarCalendar.Date, own.Openid, formId)

		// 更新formId
		err = item.UpdateFromID(time.Now().Format("20060102"))
		if err != nil {
			fmt.Println("updatesend sms date err: ", err)
		}

	}

	fmt.Println("list", list)
	return

}

func Notify() {
	fmt.Println("进入Notify")
	for i := 0; i < 2; i++ {
		users, err := bmob.UserList()
		if err != nil {
			fmt.Println("user list err: ", err)
			continue
		}
		list, err := bmob.BirthdayList(nil)
		if err != nil {
			fmt.Println("birtday list err: ", err)
			continue
		}
		fmt.Println("进入For", i)
		// fmt.Println("list",list)
		mUser := bmob.UserListToMap(users)
		for _, item := range list {

			birthday, err := time.ParseInLocation("2006-01-02 15:04:05", item.SolarCalendar.Date, time.Now().Location())
			if err != nil {
				fmt.Println("parse in location: ", err)
				continue
			}
			//fmt.Println("birthday: ", birthday.Day(), birthday.Month(), )
			//
			birthdayYmd := birthday.Format("0102")
			//日系日期
			reminderDate := time.Now().Add(48 * time.Hour).Format("0102")
			if item.Own == 13 { //调试信息
				fmt.Println("item", item)
				fmt.Println("birthday.Format", birthdayYmd, reminderDate)
			}
			// if birthdayYmd == reminderDate && item.Own == 13 {
			if birthdayYmd == reminderDate {
				var own *bmob.User
				if v, ok := mUser[int(item.Own)]; ok {
					own = v
				} else {
					continue
				}
				//phone := strconv.FormatFloat(own.Phone, 'f', -1, 64)
				//if own.Phone == "18078867423" {

				fmt.Println("own", own)

				if item.SendSmsDate != time.Now().Format("20060102") { //判断是否发过短信
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
					// err = weichatNotify(own.UserData.Openid, "亲爱的甩甩，后天是你女票生日，记得发礼物喔", time.Now().Add(48*time.Hour).Format("2006-01-02"))
					// if err != nil {
					// 	fmt.Println("weichatnotify err: ", err)
					// }
				}
				//}

			}
		}
		time.Sleep(time.Second * 60)
	}
}

func weichatNotify(openid string, note string, birthdayTime string, formId string) error {
	token, err := getWeiChatToken()
	fmt.Println("token", token)
	if err != nil {
		return err
	}
	templateId := models.AppConfig.DefaultString("weichat::template_id", "")
	urls := "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=" + token
	m := map[string]interface{}{
		"touser":      openid,
		"template_id": templateId,
		"page":        "",
		"form_id":     formId,
		"data": map[string]interface{}{
			"keyword1": map[string]string{
				"value": "生日提醒",
				"color": "#173177",
			},
			"keyword2": map[string]string{
				"value": note,
				"color": "#173177",
			},
			"keyword3": map[string]string{
				"value": birthdayTime,
				"color": "#173177",
			},
		},
	}
	b, err := request.WPost(urls, m)
	if err != nil {
		return err
	}
	fmt.Println("wchat post res: ", string(b))
	return nil
}

func getWeiChatToken() (string, error) {
	appid := models.AppConfig.DefaultString("weichat::appid", "")
	secret := models.AppConfig.DefaultString("weichat::secret", "")
	m := map[string]string{
		"grant_type": "client_credential",
		"appid":      appid,
		"secret":     secret,
	}
	address := models.AppConfig.DefaultString("weichat::token_url", "")
	b, err := request.WGet(address, m)
	if err != nil {
		return "", err
	}
	var res map[string]interface{}
	if err = json.Unmarshal(b, &res); err != nil {
		return "", err
	}
	token := ""
	if v, ok := res["access_token"].(string); ok {
		token = v
	}
	if len(token) <= 0 {
		return "", errors.New("invalid token.")
	}
	return token, nil
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
