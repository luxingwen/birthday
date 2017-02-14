package bmob

var address map[string]string

func init() {
	address = map[string]string{
		"birthday_list": "https://api.bmob.cn/1/classes/birthday_list",
		"user":          "https://api.bmob.cn/1/classes/user",
	}
}

func getAddress(m string) string {
	if v, ok := address[m]; ok {
		return v
	}
	return ""
}
