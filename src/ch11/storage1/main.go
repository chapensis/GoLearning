package main

import "fmt"

func bytesInuse(username string) int64 {
	return 1 /* ... */
}
// 邮件发送者配置
// 注意：永远不要把密码放到源码中
const sender = "notifications@example.com"
const password = "123456"
const hostname = "smtp.example.com"

const template = `Warning you are using %d bytes of storage, %d%% of your quota.`

func CheckQuota(username string) {
	used := bytesInuse(username)
	const quota = 98
	percent := quota * used
	if percent < 90 {
		fmt.Println("percent:", percent)
		return
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}

var notifyUser = func (username string, msg string) {

}
