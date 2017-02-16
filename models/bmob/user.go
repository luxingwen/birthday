package bmob

import (
	"birthday/request"
	"encoding/json"
	"fmt"
	"strconv"
)

//{"expires_in":7200,"openid":"ocBzs0DbpkQNpg6SsrCf7rhze8GY","session_key":"w0cRl0e1y49i9o0YQsxM8w=="}
type User struct {
	UserName string  `json:"username"`
	Phone    string  `json:"mobilePhoneNumber"`
	Uid      float64 `json:"uid"`
	UserData struct {
		ExpiresIn  float64 `json:"expires_in"`
		Openid     string  `json:"openid"`
		SessionKey string  `json:"session_key"`
	} `json:"userData"`
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
			fmt.Println("usersOnpage on page error: ", err)
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
