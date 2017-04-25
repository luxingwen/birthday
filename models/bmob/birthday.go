package bmob

import (
	"birthday/request"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Birthday struct {
	ObjectId      string  `json:"objectId"`
	Sex           bool    `json:"sex"`
	Name          string  `json:"name"`
	Own           int64   `json:"own"`
	Phone         float64 `json:"phone"`
	SendSmsDate   string  `json:"sendSmsDate"`
	FormId        string  `json:"formId"`
	SolarCalendar struct {
		Type string `json:"type"`
		Date string `json:"iso"`
	} `json:"solarCalendar"`
	CreatedAt string `json:"createdAt"`
}

func count(address string, where map[string]string) (int, error) {
	m := map[string]string{"limit": "0", "count": "1"}
	if len(where["createdAt"]) > 0 {
		// m["where"] = "{\"formId\":{\"$ne\":\"\"},\"createdAt\":{\"$gte\":{\"__type\": \"Date\", \"iso\": \"2017-04-23 23:59:59\"}}}"
		reminderDate := time.Now().AddDate(0, 0, -3).Format("2006-01-02 15:04:05")
		whereStr := fmt.Sprintf("{\"formId\":{\"$ne\":\"\"},\"createdAt\":{\"$gte\":{\"__type\": \"Date\", \"iso\": \"%s\"}}}", reminderDate)
		m["where"] = whereStr
		fmt.Println("whereStr", whereStr)
	}
	b, err := request.Get(address, m)
	if err != nil {
		return 0, err
	}
	var res map[string]interface{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return 0, err
	}
	if v, ok := res["count"].(float64); ok {
		return int(v), nil
	}
	return 0, nil
}

func BirthdayList(where map[string]string) ([]*Birthday, error) {
	count, err := count(getAddress("birthday_list"), where)
	if err != nil {
		return nil, err
	}
	var r []*Birthday
	for i := 0; i < count; i += 100 {

		v, err := BirthdayOnPage("100", strconv.Itoa(i), where)
		if err != nil {
			fmt.Println("birthday on page error: ", err)
			continue
		}
		r = append(r, v...)
		fmt.Println("rrr", r, count, v)
	}
	return r, nil
}

func BirthdayOnPage(limit, skip string, where map[string]string) ([]*Birthday, error) {
	m := map[string]string{"limit": limit, "skip": skip}

	if len(where["createdAt"]) > 0 {
		reminderDate := time.Now().AddDate(0, 0, -3).Format("2006-01-02 15:04:05")
		whereStr := fmt.Sprintf("{\"formId\":{\"$ne\":\"\"},\"createdAt\":{\"$gte\":{\"__type\": \"Date\", \"iso\": \"%s\"}}}", reminderDate)
		m["where"] = whereStr
	}
	b, err := request.Get(getAddress("birthday_list"), m)
	if err != nil {
		return nil, err
	}
	var res struct {
		Data []*Birthday `json:"results"`
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (this *Birthday) UpdateSendSmsDate(date string) error {
	m := map[string]string{"sendSmsDate": date}
	r, err := request.Put(getAddress("birthday_list")+"/"+this.ObjectId, m)
	if err != nil {
		return err
	}
	fmt.Println(string(r))
	return nil
}

func (this *Birthday) UpdateFromID(date string) error {
	m := map[string]string{"formId": ""}
	r, err := request.Put(getAddress("birthday_list")+"/"+this.ObjectId, m)
	if err != nil {
		return err
	}
	fmt.Println(string(r))
	return nil
}
