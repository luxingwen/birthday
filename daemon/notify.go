package daemon

import (
	"birthday/models/bmob"
	"birthday/request"
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
				continue
			}
			if birthday.Month() == time.Now().Month() && birthday.Day() == time.Now().Day() {
				var own *bmob.User
				if v, ok := mUser[int(item.Own)]; ok {
					own = v
				} else {
					continue
				}
				//phone := strconv.FormatFloat(own.Phone, 'f', -1, 64)

				if item.SendSmsDate != time.Now().Format("20060102") {
					if own.Phone == "18078867423" {
						r, err := sendSms(own.Phone, "亲爱的甩甩，情人节快乐")
						if err != nil {
							fmt.Println("send sms err:", err)
							continue
						}
						fmt.Println(string(r))
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
		"content":           content,
	}
	r, err = request.Post("https://api.bmob.cn/1/requestSms", data)
	if err != nil {
		return
	}
	return
}
