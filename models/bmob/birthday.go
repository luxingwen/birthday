package bmob

import (
	"birthday/request"
	"encoding/json"
	"fmt"
	"strconv"
)

type Birthday struct {
	ObjectId      string  `json:"objectId"`
	Sex           bool    `json:"sex"`
	Name          string  `json:"name"`
	Own           int64   `json:"own"`
	Phone         float64 `json:"phone"`
	SendSmsDate   string  `json:"sendSmsDate"`
	SolarCalendar struct {
		Type string `json:"type"`
		Date string `json:"iso"`
	} `json:"solarCalendar"`
}

func count(address string) (int, error) {
	m := map[string]string{"limit": "0", "count": "1"}
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

func BirthdayList() ([]*Birthday, error) {
	count, err := count(getAddress("birthday_list"))
	if err != nil {
		return nil, err
	}
	var r []*Birthday
	for i := 0; i < count; i += 100 {

		v, err := BirthdayOnPage("100", strconv.Itoa(i))
		if err != nil {
			fmt.Println("birthday on page error: ", err)
			continue
		}
		r = append(r, v...)
	}
	return r, nil
}

func BirthdayOnPage(limit, skip string) ([]*Birthday, error) {
	m := map[string]string{"limit": limit, "skip": skip}
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
