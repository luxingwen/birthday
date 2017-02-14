package bmob

import (
	"birthday/request"
	"encoding/json"
	"fmt"
	"strconv"
)

type User struct {
	UserName string  `json:"username"`
	Phone    string  `json:"mobilePhoneNumber"`
	Uid      float64 `json:"uid"`
}

func UserList() ([]*User, error) {
	count, err := count(getAddress("user"))
	if err != nil {
		return nil, err
	}
	var r []*User
	for i := 0; i < count; i += 100 {
		v, err := UsersOnPage("100", strconv.Itoa(i))
		if err != nil {
			fmt.Println("birthday on page error: ", err)
			continue
		}
		r = append(r, v...)
	}
	return r, nil
}

func UsersOnPage(limit, skip string) ([]*User, error) {
	m := map[string]string{"limit": limit, "skip": skip}
	b, err := request.Get(getAddress("user"), m)
	if err != nil {
		return nil, err
	}
	var res struct {
		Data []*User `json:"results"`
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func UserListToMap(r []*User) map[int]*User {
	m := make(map[int]*User)
	for _, v := range r {
		m[int(v.Uid)] = v
	}
	return m
}
