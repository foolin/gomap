"# gomap" 

type StrMap map[string]interface{}

include get base type method.

```go

	jstr := `
	{"userId":123,"name":"Forin", "sex": "2", "createTime":"2016-06-14T14:54:18.6492677+08:00","ext":{"son":"Tigo","Age":2}}
	`
	user := gomap.NewStrMap()
	err := json.Unmarshal([]byte(jstr), &user)
	if err != nil {
		log.Printf("json error: %v", err)
		return
	}

	log.Printf("\n===================\n%#v\n ", user)
	log.Printf("\n===================\nUserId：%v\n ", user.Int("userId", 0))

```