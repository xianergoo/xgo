package main

type Data struct {
	Ip       string   `json:"ip"`
	User     stirng   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
}
