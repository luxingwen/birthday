package bmob

var address map[string]string

func init() {
	address = map[string]string{
		"birthday_list": "https://api.bmob.cn/1/classes/birthday_list",
		"user":          "https://api.bmob.cn/1/users",
		"shop":          "https://api.bmob.cn/1/classes/shop",
		"shopcake2":     "https://api.bmob.cn/1/classes/shopCake2",
	}
}

func getAddress(m string) string {
	if v, ok := address[m]; ok {
		return v
	}
	return ""
}
